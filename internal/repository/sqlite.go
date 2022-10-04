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
	//Session table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Sessions(
		user_id INTEGER NOT NULL,
		uuid TEXT NOT NULL,
		expires DATETIME NOT NULL,
		FOREIGN KEY(user_id) REFERENCES User(id) ON DELETE CASCADE
	)`); err != nil {
		return err
	}

	//User table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS User(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL, 
		nickname TEXT UNIQUE NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		password TEXT NOT NULL,
		gender Text NOT NULL,
		age TEXT NOT NULL
	)
	`); err != nil {
		return err
	}

	//Post table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Posts(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		tags TEXT NOT NULL,
		content TEXT NOT NULL,
		FOREIGN KEY(user_id) REFERENCES User(id) ON DELETE CASCADE
	)`); err != nil {
		return err
	}

	//Comment table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Comments(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		post_id INTEGER NOT NULL,
		commenter_id INTEGER NOT NULL,
		FOREIGN KEY(commenter_id) REFERENCES User(id) ON DELETE CASCADE,
		FOREIGN KEY(post_id) REFERENCES Posts(id) ON DELETE CASCADE
	)
	`); err != nil {
		return err
	}

	// Post and Comment rating table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS PostRating(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		likes int,
		post_id INTEGER NOT NULL,
		liked_userId INTEGER NOT NULL,
		FOREIGN KEY(post_id) REFERENCES Posts(id) ON DELETE CASCADE,
		FOREIGN KEY(liked_userId) REFERENCES User(id) ON DELETE CASCADE
	)`); err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS CommentRating(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		likes int,
		comment_id INTEGER NOT NULL,
		liked_userId INTEGER NOT NULL,
		FOREIGN KEY(comment_id) REFERENCES Comments(id) ON DELETE CASCADE,
		FOREIGN KEY(liked_userId) REFERENCES User(id) ON DELETE CASCADE
	)`); err != nil {
		return err
	}

	return nil
}
