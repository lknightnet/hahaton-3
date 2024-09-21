package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

func New(ctx context.Context, connURI string) (*Postgres, error) {
	pg := &Postgres{}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	cfg, err := pgxpool.ParseConfig(connURI)
	if err != nil {
		return nil, fmt.Errorf("postgres/new: parse connURI is failed: %s", err.Error())
	}

	cfg.MaxConns = 10

	pg.Pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres/new: connect with postgres is failed: %s", err.Error())
	}

	err = pg.ping(ctx, 10)
	if err == nil {
		return pg, nil
	}

	return nil, err
}

func (p *Postgres) ping(ctx context.Context, retriesLeft int) error {
	err := p.Pool.Ping(ctx)
	if err != nil {
		fmt.Printf("Error pinging PostgreSQL: %v\n", err)

		if retriesLeft > 0 {
			time.Sleep(time.Second) // Можете использовать другую задержку
			return p.ping(ctx, retriesLeft-1)
		}

		return fmt.Errorf("failed to ping PostgreSQL after retries")
	}
	fmt.Println("Теперь работает")
	return nil
}
