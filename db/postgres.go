package db

import (
	"context"
	"fmt"
	"log"
	//"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Psql struct {
	Pool *pgxpool.Pool
}

func PsqlNewClient(ctx context.Context, dsn string) (*Psql, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse DSN failed: %w", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 1
	cfg.MaxConnLifetime = time.Hour

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("connect to DB failed: %w", err)
	}

	log.Println("✅ Connected to PostgreSQL")
	return &Psql{Pool: pool}, nil
}

// Close gracefully closes the pool
func (p *Psql) Close() {
	p.Pool.Close()
	log.Println("✅ PostgreSQL connection closed")
}

