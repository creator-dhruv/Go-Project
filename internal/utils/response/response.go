package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status int
	Error  string
}

func WriteJson(w http.ResponseWriter, status int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(status int, err error) Response {
	return Response{
		Status: status,
		Error:  err.Error(),
	}
}

func ValidationError(error validator.ValidationErrors) Response {
	var errMsg []string

	for _, err := range error {
		switch err.ActualTag() {
		case "required":
			errMsg = append(errMsg, fmt.Sprintf("%s is required field", err.Field()))
		default:
			errMsg = append(errMsg, fmt.Sprintf("%s is invalid", err.Field()))
		}
	}

	return Response{
		Error:  errMsg[0],
		Status: 401,
	}
}
