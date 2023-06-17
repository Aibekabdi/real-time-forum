package postgres

import (
	"fmt"
	"forum/pkg/utils"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg utils.Database) (*sqlx.DB, error) {
	// Connect to the PostgreSQL database
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DBName,
		cfg.Password,
		cfg.SSLMode,
	))
	if err != nil {
		return nil, err
	}
	// Ping the database to ensure a connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
