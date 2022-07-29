package db

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewSQLStore(db *sql.DB) *SQLStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
