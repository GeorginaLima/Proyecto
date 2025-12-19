package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

// InitDBFromEnv initializes a CockroachDB/Postgres-compatible connection using
// the `COCKROACH_DSN` environment variable. If unset, it falls back to a
// local default. It also runs a minimal schema migration for `reservations`.
func InitDBFromEnv() error {
	dsn := os.Getenv("COCKROACH_DSN")
	if dsn == "" {
		dsn = "postgresql://root@localhost:26257/defaultdb?sslmode=disable"
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(5)
	db.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("ping db: %w", err)
	}

	DB = db

	if err := ensureSchema(ctx, db); err != nil {
		return fmt.Errorf("ensure schema: %w", err)
	}

	return nil
}

func CloseDB() error {
	if DB == nil {
		return nil
	}
	return DB.Close()
}

func ensureSchema(ctx context.Context, db *sql.DB) error {
	ddl := `CREATE TABLE IF NOT EXISTS reservations (
		id TEXT PRIMARY KEY,
		customer_name TEXT,
		date TIMESTAMPTZ,
		status TEXT,
		created_at TIMESTAMPTZ DEFAULT now()
	);`
	_, err := db.ExecContext(ctx, ddl)
	return err
}
