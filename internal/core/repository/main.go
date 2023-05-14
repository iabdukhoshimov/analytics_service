package repository

import (
	"context"

	"gitlab.com/greatsoft/xif-backend/internal/config"
	"gitlab.com/greatsoft/xif-backend/internal/core/repository/psql"
	"gitlab.com/greatsoft/xif-backend/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, cfg *config.Config) Store {
	return psql.NewStore(ctx, cfg.PSQL.URI)
}
