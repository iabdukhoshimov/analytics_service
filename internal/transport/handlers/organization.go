package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voxe-analytics/internal/core/domain"
	"gitlab.com/voxe-analytics/pkg/processor"
)

// @Summary Get all organizations
// @Tags organizations
// @Param filter query domain.GetAllParams true "Organization ID"
// @Success 200 {object} domain.OrganizationGetAll
// @Failure 400 {object} response.Error
// @Router /organizations [get]
// @Security ApiKeyAuth
func (h *handler) getOrganizations(c *gin.Context) {
	var (
		filter domain.GetAllParams
		resp   domain.OrganizationGetAll
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.handleError(c, err)
		return
	}

	dbRes, err := processor.ExecuteWithResp(ctx, filter, h.repo.OrganizationGetAll)
	if h.handleError(c, err) {
		return
	}

	count, err := processor.ExecuteWithResp(ctx, filter, h.repo.OrganizationGetAllCount)
	if h.handleError(c, err) {
		return
	}

	resp.Objects = dbRes
	resp.Count = count

	c.JSON(http.StatusOK, resp)
}
