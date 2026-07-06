package sqlite

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/creator-dhruv/Go-Project/internal/config"
	"github.com/creator-dhruv/Go-Project/internal/types"
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

func (s *Sqlite) GetUserById(id int64) (types.User, error) {
	stmt, err := s.Db.Prepare(`SELECT id, name, email, age, created_at, updated_at FROM user WHERE id = ? LIMIT 1`)
	if err != nil {
		return types.User{}, err
	}

	defer stmt.Close()

	var user types.User

	err = stmt.QueryRow(id).Scan(&user.Id, &user.Name, &user.Email, &user.Age, &user.Created_At, &user.Updated_At)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.User{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.User{}, err
	}

	return user, nil
}
