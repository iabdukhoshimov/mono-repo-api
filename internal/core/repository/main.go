package repository

import (
	"context"

	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql"
	"github.com/abdukhashimov/go_api_mono_repo/internal/core/repository/psql/sqlc"
)

type Store interface {
	sqlc.Querier
}

func New(ctx context.Context, dsn string) Store {
	return psql.NewStore(ctx, dsn)
}
