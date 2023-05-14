package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/voxe-analytics/internal/core/domain"
	"gitlab.com/voxe-analytics/pkg/processor"
)

// @Summary Get all licenses
// @Tags licenses
// @Param filter query domain.LicenseFilter true "License ID"
// @Success 200 {object} domain.LicenseGetAll
// @Failure 400 {object} response.Error
// @Router /licenses [get]
func (h *handler) getAllLicenses(c *gin.Context) {
	var (
		params domain.LicenseFilter
		resp   domain.LicenseGetAll
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindQuery(&params); err != nil {
		h.handleError(c, err)
		return
	}

	dbRes, err := processor.ExecuteWithResp(ctx, params, h.repo.LicenseGetAll)
	if h.handleError(c, err) {
		return
	}

	count, err := processor.ExecuteWithResp(ctx, params, h.repo.LicenseGetAllCount)
	if h.handleError(c, err) {
		return
	}

	resp.Objects = dbRes
	resp.Count = count

	c.JSON(200, resp)
}
