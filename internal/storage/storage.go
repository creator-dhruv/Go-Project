package storage

import (
	"time"

	"github.com/creator-dhruv/Go-Project/internal/types"
)

type Storage interface {
	CreateUser(name string, email string, age int, created_at time.Time) (int64, error)
	GetUserById(id int64) (types.User, error)
	GetUsers() ([]types.User, error)
}
