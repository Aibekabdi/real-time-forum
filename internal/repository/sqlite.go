package repository

import (
	"database/sql"
	"path/filepath"

	"forum/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(config config.Database) (*sql.DB, error) {
	db, err := sql.Open(config.Driver, filepath.Join(config.Path, config.FileName))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	if err := createTables(db); err != nil {
		return nil, err
	}
	return db, nil
}

func createTables(db *sql.DB) error {
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS User(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL, 
		username TEXT UNIQUE NOT NULL, 
		password TEXT NOT NULL
	)
	`); err != nil {
		return err
	}
	return nil
}
