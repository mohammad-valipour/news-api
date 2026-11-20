package database

import (
	"database/sql"
	"fmt"
	"news/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Connect connects to PostgreSQL
func Connect(cfg *config.Config) (*sql.DB, error) {

	// Create database connection string
	// DSN = Data Source Name
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	// Open database connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Check database connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
