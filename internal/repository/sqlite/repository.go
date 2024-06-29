package sqlite

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewSQLiteRepository(
	log *slog.Logger,
	db *sql.DB,
) *SQLiteRepository {
	return &SQLiteRepository{db: db, log: log}
}
