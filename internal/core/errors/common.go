package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorApp struct {
	HttpCode int
	Code     string `json:"code"`
	Message  string `json:"message"`
}

func (e *ErrorApp) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func IsErrorApp(err error) (*ErrorApp, bool) {
	var appErr *ErrorApp
	if errors.As(err, &appErr) {
		return appErr, true
	}
	return nil, false
}

func RepetitionRule_Error() *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_REPETITION_RULE",
		Message:  "the repetition rule is not specified",
	}
}

func StartDate_Error() *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_START_DATE",
		Message:  "incorrect start date",
	}
}

func Interval_D_Error(msg string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_INTERVAL_FOR_D",
		Message:  msg,
	}
}

func Interval_W_Error(msg string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_INTERVAL_FOR_W",
		Message:  msg,
	}
}

func Interval_M_Error(msg string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_INTERVAL_FOR_M",
		Message:  msg,
	}
}

func Rule_Error(invalidRule string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "INVALID_RULE",
		Message:  fmt.Sprintf("unsupported rule: %s", invalidRule),
	}
}

func BadRequest(err string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusBadRequest,
		Code:     "BAD_REQUEST",
		Message:  err,
	}
}

func ServerError(err string) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusInternalServerError,
		Code:     "INTERNAL_SERVER_ERROR",
		Message:  err,
	}
}

func TaskNotFaund(id int) *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusNotFound,
		Code:     "TASK_NOT_FAUND",
		Message:  fmt.Sprintf("task with id %d not faund", id),
	}
}

func InvalidJWT() *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusUnauthorized,
		Code:     "JWT_NOT_VALID",
		Message:  "JWT-token not valid",
	}
}

func InvalidPassword() *ErrorApp {
	return &ErrorApp{
		HttpCode: http.StatusUnauthorized,
		Code:     "PASSWORD_NOT_VALID",
		Message:  "password not valid",
	}
}
