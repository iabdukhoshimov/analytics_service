package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voxe-analytics/internal/core/domain"
	"gitlab.com/voxe-analytics/pkg/jwt"
	"gitlab.com/voxe-analytics/pkg/processor"
	"golang.org/x/crypto/bcrypt"
)

// @Summary login with email
// @Tags auth
// @Param auth body domain.Login true "response.Success"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Router /login [post]
func (h *handler) login(c *gin.Context) {
	payload := domain.Login{}
	err := c.ShouldBindJSON(&payload)
	if h.handleError(c, err) {
		return
	}

	ctx, cancel := h.makeContext()
	defer cancel()

	dbUser, err := processor.ExecuteWithResp(ctx, payload.Email, h.repo.UserGetOneByEmail)
	if h.handleError(c, err) {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), []byte(payload.HashedPassword))
	if err != nil {
		h.handleError(c, errorWrongCredentials)
		return
	}

	jwtMap := map[string]interface{}{
		"id":   dbUser.ID,
		"role": dbUser.RoleID,
	}

	token, err := h.jwt.GenerateJWT(jwtMap, jwt.AccessToken)
	if h.handleError(c, err) {
		return
	}

	refToken, err := h.jwt.GenerateJWT(jwtMap, jwt.RefreshToken)
	if h.handleError(c, err) {
		return
	}

	resp := domain.LoginResponse{}
	resp.AccessToken = token
	resp.RefreshToken = refToken
	resp.AccessExpiresAt = h.config.JWT.AccessTokenTTLMinutes.Microseconds()
	resp.RefreshExpiresAt = h.config.JWT.RefreshTokenTTLHours.Microseconds()

	c.JSON(200, resp)
}

// @Summary Get user by auth token
// @Tags auth
// @Success 200 {object} sqlc.User
// @Failure 400 {object} response.Error
// @Router /auth/me [get]
// @Security ApiKeyAuth
func (h *handler) getUserByToken(c *gin.Context) {
	userID := c.GetString("id")
	if userID == "" {
		h.handleError(c, errorWrongCredentials)
		return
	}

	ctx, cancel := h.makeContext()
	defer cancel()

	dbRes, err := processor.ExecuteWithResp(ctx, userID, h.repo.UserGetOne)
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, dbRes)
}
