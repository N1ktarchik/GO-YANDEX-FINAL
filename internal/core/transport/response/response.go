package response

import (
	"encoding/json"
	"n1ktarchik/go-final/internal/core/errors"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

type JWTResponse struct {
	Token string `json:"token"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"code":500, "error":"INTERNAL_SERVER_ERROR"}`))
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(resp)
}

func RespondWithJWT(w http.ResponseWriter, code int, jwt string) {
	RespondWithJSON(w, code, JWTResponse{Token: jwt})
}

func RespondWithVallue(w http.ResponseWriter, code int, key string, data any) {
	RespondWithJSON(w, code, map[string]any{
		key: data,
	})
}

func RespondWithError(w http.ResponseWriter, err error) {
	if appErr, ok := errors.IsErrorApp(err); ok {
		RespondWithJSON(w, appErr.HttpCode, ErrorResponse{
			Error:   appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	RespondWithJSON(w, http.StatusInternalServerError, ErrorResponse{
		Error:   "INTERNAL_SERVER_ERROR",
		Message: "Something went wrong",
	})
}
