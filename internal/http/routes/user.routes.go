package routes

import (
	"net/http"

	"github.com/creator-dhruv/Go-Project/internal/http/handlers/user"
)

func UserRouter(router *http.ServeMux) {
	router.HandleFunc("GET /", user.New())
}
