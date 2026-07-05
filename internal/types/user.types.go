package types

import "time"

type User struct {
	Id            int64     `json:"id"`
	Name          string    `json:"name" validate:"required"`
	Email         string    `json:"email" validate:"required"`
	Age           int       `json:"age" validate:"required"`
	Created_At    time.Time `json:"created_at"`
	Updated_At    time.Time `json:"updated_at"`
	Refresh_Token string    `json:"refresh_token"`
}
