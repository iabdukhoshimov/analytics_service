package domain

import "gitlab.com/voxe-analytics/internal/core/repository/psql/sqlc"

type PaymentGetAll struct {
	Objects []sqlc.Payment `json:"objects"`
	Count   int64          `json:"count"`
}

type PaymentFilter struct {
	Limit          int32  `form:"limit" json:"limit" default:"10"`
	Offset         int32  `form:"offset" json:"offset"`
	Status         int32  `form:"status" json:"status"`
	Type           int32  `form:"type" json:"type"`
	OrganizationID string `form:"organization_id" json:"organization_id"`
}
