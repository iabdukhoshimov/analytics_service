package repository

import (
	"context"

	"gitlab.com/voxe-analytics/internal/config"
	"gitlab.com/voxe-analytics/internal/core/repository/psql"
	"gitlab.com/voxe-analytics/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, cfg *config.Config) Store {
	return psql.NewStore(ctx, cfg.PSQL.URI)
}
