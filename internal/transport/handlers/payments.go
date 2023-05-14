package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/voxe-analytics/internal/core/domain"
	"gitlab.com/voxe-analytics/pkg/processor"
)

// @Summary Get all payments
// @Tags payments
// @Param filter query domain.PaymentFilter true "Payment ID"
// @Success 200 {object} domain.PaymentGetAll
// @Failure 400 {object} response.Error
// @Router /payments [get]
// @Security ApiKeyAuth
func (h *handler) getPayments(c *gin.Context) {
	var (
		filter domain.PaymentFilter
		resp   domain.PaymentGetAll
	)

	ctx, cancel := h.makeContext()
	defer cancel()

	if err := c.ShouldBindQuery(&filter); err != nil {
		h.handleError(c, err)
		return
	}

	dbRes, err := processor.ExecuteWithResp(ctx, filter, h.repo.PaymentGetAll)
	if h.handleError(c, err) {
		return
	}

	count, err := processor.ExecuteWithResp(ctx, filter, h.repo.PaymentGetAllCount)
	if h.handleError(c, err) {
		return
	}

	resp.Objects = dbRes
	resp.Count = count

	c.JSON(http.StatusOK, resp)
}
