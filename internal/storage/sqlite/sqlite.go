package sqlite

import (
	"database/sql"
	"time"

	"github.com/creator-dhruv/Go-Project/internal/config"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(` CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT UNIQUE,
	age INTEGER,
	created_at DATE,
	updated_at DATE,
	refresh_token TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateUser(name string, email string, age int, created_at time.Time) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO user (name, email, age, created_at, updated_at, refresh_token) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age, created_at, created_at, nil)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
