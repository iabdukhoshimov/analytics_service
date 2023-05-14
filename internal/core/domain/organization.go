package domain

import "gitlab.com/voxe-analytics/internal/core/repository/psql/sqlc"

type OrganizationGetAll struct {
	Objects []sqlc.Organization `json:"objects"`
	Count   int64               `json:"count"`
}
