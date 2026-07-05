package routes

import (
	"net/http"

	"github.com/creator-dhruv/Go-Project/internal/http/handlers/user"
	"github.com/creator-dhruv/Go-Project/internal/storage"
)

func UserRouter(router *http.ServeMux, storage storage.Storage) {
	router.HandleFunc("POST /", user.New(storage))
}
