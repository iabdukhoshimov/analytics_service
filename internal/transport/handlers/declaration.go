package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voxe-analytics/internal/core/domain"
	"gitlab.com/voxe-analytics/pkg/processor"
)

// @Summary Get all declarations
// @Tags declarations
// @Param filter query domain.GetAllParams true "Declaration ID"
// @Success 200 {object} domain.DeclarationGetAll
// @Failure 400 {object} response.Error
// @Router /declarations [get]
// @Security ApiKeyAuth
func (h *handler) getAllDeclarations(c *gin.Context) {
	var (
		filter domain.GetAllParams
		resp   domain.DeclarationGetAll
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.handleError(c, err)
		return
	}

	dbRes, err := processor.ExecuteWithResp(ctx, filter, h.repo.DeclarationGetAll)
	if h.handleError(c, err) {
		return
	}

	count, err := processor.ExecuteWithResp(ctx, filter.OrganizationID, h.repo.DeclarationGetAllCount)
	if h.handleError(c, err) {
		return
	}

	resp.Objects = dbRes
	resp.Count = count

	c.JSON(http.StatusOK, resp)
}
