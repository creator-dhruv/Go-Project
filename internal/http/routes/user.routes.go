package routes

import (
	"net/http"

	"github.com/creator-dhruv/Go-Project/internal/http/handlers/user"
	"github.com/creator-dhruv/Go-Project/internal/storage"
)

func UserRouter(router *http.ServeMux, storage storage.Storage) {
	router.HandleFunc("POST /api/user/create", user.New(storage))
	router.HandleFunc("GET /api/user/{id}", user.GetById(storage))
	router.HandleFunc("GET /api/users", user.GetUsers(storage))
	router.HandleFunc("PUT /api/user/{id}", user.UpdateUser(storage))
	router.HandleFunc("DELETE /api/user/{id}", user.DeleteUser(storage))
}
