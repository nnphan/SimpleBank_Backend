package main

import (
	"context"
	"log"

	"simplebank/configs"
	"simplebank/internal/api"
	db "simplebank/internal/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    // Load config
    //cfg := configs.LoadConfig()
    // db
    //conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
    cfg, err := configs.LoadConfig("../../configs", "local")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    }

    ctx := context.Background()
    connPool, err := pgxpool.New(ctx, cfg.DB_URL)

    if err != nil {
        log.Fatal("cannot connect to DB: ", err)
    }

    
// ✅ Check connection to DB
    if err := connPool.Ping(ctx); err != nil {
        log.Fatal("cannot connect to DB: ", err)
    }

    store := db.NewSQLStore(connPool)
    server, err := api.NewServer(store)

    err = server.Start("0.0.0.0:8002")

}