package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/creator-dhruv/Go-Project/internal/storage"
	"github.com/creator-dhruv/Go-Project/internal/types"
	"github.com/creator-dhruv/Go-Project/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user types.User

		// Load data coming in body

		// Method 1 (Efficient)
		err := json.NewDecoder(r.Body).Decode(&user)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, err))
			return
		}

		// Method 2
		/*
			Byte, err := io.ReadAll(r.Body)
			if err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, err))
				return
			}

			if err := json.Unmarshal(Byte, &user); err != nil {
				response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, err))
				return
			}
		*/

		// Validate the data using validator package

		if err := validator.New().Struct(user); err != nil {
			validateErr := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErr))
			return
		}

		id, err := storage.CreateUser(user.Name, user.Email, user.Age, time.Now())
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(http.StatusInternalServerError, err))
			return
		}
		slog.Info("user created successfully")
		response.WriteJson(w, http.StatusCreated, map[string]any{"success": "OK", "userId": id})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("getting a student id : ", slog.String("id", id))

		Id, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(http.StatusBadRequest, err))
			return
		}

		user, err := storage.GetUserById(Id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(http.StatusInternalServerError, err))
			return
		}

		response.WriteJson(w, http.StatusAccepted, map[string]any{"success": "OK", "user": user})
	}
}

func GetUsers(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := storage.GetUsers()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(http.StatusInternalServerError, err))
		}

		response.WriteJson(w, http.StatusAccepted, map[string]any{"success": "OK", "users": users})

	}
}
