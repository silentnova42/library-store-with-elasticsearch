package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	client *pgxpool.Pool
}

type Config interface {
	GetConnectSting() string
}

func NewDb(ctx context.Context, config Config, attempts int) (*db, error) {
	conf, err := pgxpool.ParseConfig(config.GetConnectSting())
	if err != nil {
		return nil, err
	}

	pool, err := attemptingConnectDb(ctx, conf, attempts)
	if err != nil {
		return nil, err
	}

	return &db{
		client: pool,
	}, nil
}

func attemptingConnectDb(ctx context.Context, conf *pgxpool.Config, attempts int) (*pgxpool.Pool, error) {
	var (
		pool *pgxpool.Pool
		err  error
	)

	for attempts >= 0 {
		pool, err = pgxpool.NewWithConfig(ctx, conf)
		if err != nil {
			attempts--
			continue
		}

		if err = pool.Ping(ctx); err != nil {
			attempts--
			continue
		}

		break
	}

	return pool, nil
}
