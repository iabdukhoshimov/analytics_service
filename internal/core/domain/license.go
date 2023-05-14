package domain

import "gitlab.com/greatsoft/xif-backend/internal/core/repository/psql/sqlc"

type LicenseFilter struct {
	Limit       int32  `form:"limit" json:"limit" default:"10"`
	Offset      int32  `form:"offset" json:"offset"`
	LicenseType int32  `form:"license_type" json:"license_type"`
	Search      string `form:"search" json:"search"`
}

type LicenseGetAll struct {
	Objects []sqlc.License `json:"objects"`
	Count   int64          `json:"count"`
}
