package domain

import "gitlab.com/voxe-analytics/internal/core/repository/psql/sqlc"

type DeclarationGetAll struct {
	Objects []sqlc.DeclarationGetAllRow `json:"objects"`
	Count   int64                       `json:"count"`
}
