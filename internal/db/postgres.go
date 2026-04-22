package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(dbSource string) *pgxpool.Pool {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, dbSource)
    if err != nil {
        log.Fatalf("❌ cannot create db pool: %v", err)
    }

    // Verify connection
    if err := pool.Ping(ctx); err != nil {
        log.Fatalf("❌ cannot connect to db: %v", err)
    }

    log.Println("✅ PostgreSQL connected")
    return pool
}