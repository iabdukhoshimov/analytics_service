package domain

import "gitlab.com/greatsoft/xif-backend/internal/core/repository/psql/sqlc"

type DeclarationGetAll struct {
	Objects []sqlc.DeclarationGetAllRow `json:"objects"`
	Count   int64                       `json:"count"`
}
