package storage

import "time"

type Storage interface {
	CreateUser(name string, email string, age int, created_at time.Time) (int64, error)
}
