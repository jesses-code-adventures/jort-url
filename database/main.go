package database

import (
	"database/sql"
	"github.com/jesses-code-adventures/jort-url/passwords"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
	*passwords.PasswordHandler
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, err
	}
	encoder := passwords.NewPasswordHandler()
	database := Database{db, encoder}
	err = database.init()
	if err != nil {
		return nil, err
	}
	return &database, nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

func (db *Database) init() error {
	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table'`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return err
		}
		if name == "user" || name == "url" {
			return nil
		}
	}
	return db.createTables()
}

func (db *Database) createTables() error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		session_jwt TEXT
	) STRICT`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS url (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		created_at TEXT NOT NULL,
		user_id INT NOT NULL,
		url TEXT NOT NULL,
		short_pathname TEXT NOT NULL,
		clicks INT NOT NULL DEFAULT 0,
		FOREIGN KEY (user_id) REFERENCES user(id),
		UNIQUE(url, user_id)
	) STRICT`)
	if err != nil {
		return err
	}
	return nil
}
