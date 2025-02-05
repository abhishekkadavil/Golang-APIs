package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-lang-test-apis/internal/types"
	"go-lang-test-apis/internal/utils/response"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
)

/**
 * @author Abhishek Kadavil
 */

func NewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Creating user")

		var gouser types.GoApiUser

		err := json.NewDecoder(r.Body).Decode(&gouser)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusUnprocessableEntity, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		// Mandatory filed validation validation
		if err := validator.New().Struct(gouser); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusUnprocessableEntity, response.ValidationError(validateErrs))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"status": "OK"})
	}
}
