package handlers

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get all license types
// @Tags types
// @Success 200 {object} []sqlc.LicenseType
// @Failure 400 {object} response.Error
// @Router /types/licenses [get]
func (h *handler) getLicenseTypes(c *gin.Context) {
	ctx, cancel := h.makeContext()
	defer cancel()

	dbRes, err := h.repo.LicenseTypesGetAll(ctx)
	if h.handleError(c, err) {
		return
	}

	c.JSON(200, dbRes)
}

// @Summary Get all statuses
// @Tags types
// @Success 200 {object} []sqlc.Status
// @Failure 400 {object} response.Error
// @Router /types/statuses [get]
func (h *handler) getStatuses(c *gin.Context) {
	ctx, cancel := h.makeContext()
	defer cancel()

	dbRes, err := h.repo.StatusesGetAll(ctx)
	if h.handleError(c, err) {
		return
	}

	c.JSON(200, dbRes)
}

// @Summary Get all payment types
// @Tags types
// @Success 200 {object} []sqlc.PaymentType
// @Failure 400 {object} response.Error
// @Router /types/payments [get]
func (h *handler) getPaymentTypes(c *gin.Context) {
	ctx, cancel := h.makeContext()
	defer cancel()

	dbRes, err := h.repo.PaymentTypesGetAll(ctx)
	if h.handleError(c, err) {
		return
	}

	c.JSON(200, dbRes)
}
