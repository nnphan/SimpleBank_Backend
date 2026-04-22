package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

//Store provides all functions to execute db queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewSQLStore(connPool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

