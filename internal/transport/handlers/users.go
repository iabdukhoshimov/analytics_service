package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"

	"gitlab.com/greatsoft/xif-backend/internal/core/domain"
	"gitlab.com/greatsoft/xif-backend/internal/pkg/response"
	"gitlab.com/greatsoft/xif-backend/pkg/processor"
)

// @Summary Create user
// @Tags users
// @Param user body domain.UserCreateParams true "response.Success"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Router /users [post]
// @Security ApiKeyAuth
func (h *handler) createUser(c *gin.Context) {
	var (
		userParams domain.UserCreateParams
		resp       response.Success
	)

	if err := c.ShouldBindJSON(&userParams); err != nil {
		h.handleError(c, err)
		return
	}

	ctx, cancel := h.makeContext()
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParams.HashedPassword), bcrypt.DefaultCost)
	if h.handleError(c, err) {
		return
	}

	userParams.HashedPassword = string(hashedPassword)

	dbRes, err := processor.ExecuteWithResp(ctx, userParams, h.repo.UserInsertOne)
	if h.handleError(c, err) {
		return
	}

	resp.ID = dbRes
	c.JSON(http.StatusOK, resp)
}

// @Summary Get user by id
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} sqlc.User
// @Failure 400 {object} response.Error
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func (h *handler) getUserByID(c *gin.Context) {
	var (
		userID = c.Param("id")
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	dbRes, err := processor.ExecuteWithResp(ctx, userID, h.repo.UserGetOne)
	if h.handleError(c, err) {
		return
	}

	c.JSON(http.StatusOK, dbRes)
}

// @Summary Update user by id
// @Tags users
// @Param id path string true "User ID"
// @Param user body domain.UserCreateParams true "domain.UserCreateParams"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func (h *handler) updateUserByID(c *gin.Context) {
	var (
		userID  = c.Param("id")
		payload domain.UserCreateParams
		resp    response.Success
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindJSON(&payload); err != nil {
		h.handleError(c, err)
		return
	}

	if payload.HashedPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.HashedPassword), bcrypt.DefaultCost)
		if h.handleError(c, err) {
			return
		}

		payload.HashedPassword = string(hashedPassword)
	}

	payload.ID = userID
	if err := processor.Execute(ctx, payload, h.repo.UserUpdateOne); err != nil {
		h.handleError(c, err)
		return
	}

	resp.ID = userID
	c.JSON(http.StatusOK, resp)
}

// @Summary Delete user by id
// @Tags users
// @Param id path string true "User ID"
// @Success 200 {object} response.Success
// @Failure 400 {object} response.Error
// @Router /users/{id} [delete]
// @Security ApiKeyAuth
func (h *handler) deleteUserByID(c *gin.Context) {
	var (
		userID = c.Param("id")
		resp   response.Success
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := processor.Execute(ctx, userID, h.repo.UserDeleteOne); err != nil {
		h.handleError(c, err)
		return
	}

	resp.ID = userID
	c.JSON(http.StatusOK, resp)
}

// @Summary Get users
// @Tags users
// @Param filter query domain.GetAllParams false "filter"
// @Success 200 {object} domain.Users
// @Failure 400 {object} response.Error
// @Router /users [get]
// @Security ApiKeyAuth
func (h *handler) getUsers(c *gin.Context) {
	var (
		resp    domain.Users
		filters domain.GetAllParams
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindQuery(&filters); err != nil {
		h.handleError(c, err)
		return
	}

	dbRes, err := processor.ExecuteWithResp(ctx, filters, h.repo.UserGetAll)
	if h.handleError(c, err) {
		return
	}

	count, err := processor.ExecuteWithResp(ctx, filters, h.repo.UserGetAllCount)
	if h.handleError(c, err) {
		return
	}

	resp.Objects = dbRes
	resp.Count = int32(count)

	c.JSON(http.StatusOK, resp)
}
