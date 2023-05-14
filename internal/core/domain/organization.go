package domain

import "gitlab.com/greatsoft/xif-backend/internal/core/repository/psql/sqlc"

type OrganizationGetAll struct {
	Objects []sqlc.Organization `json:"objects"`
	Count   int64               `json:"count"`
}

