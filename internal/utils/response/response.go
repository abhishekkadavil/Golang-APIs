package response

import (
	"encoding/json"
	"net/http"
)

/**
 * @author Abhishek Kadavil
 */

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

const (
	StatusOK    = "OK"
	StatusError = "NOK"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
	return Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}
