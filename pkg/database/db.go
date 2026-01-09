package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDbPool(db_dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), db_dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func CloseDB(conn *pgxpool.Pool) {
	conn.Close()
}
