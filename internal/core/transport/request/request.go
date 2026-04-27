package request

import (
	core_errors "n1ktarchik/go-final/internal/core/errors"

	"encoding/json"
	"io"
	"net/http"
)

func DecodeJSON(r *http.Request, userData any) error {
	reqData, err := io.ReadAll(r.Body)
	if err != nil {
		return core_errors.ServerError("error to read request body")
	}

	defer r.Body.Close()

	if len(reqData) == 0 {
		return core_errors.BadRequest("length of request data can not be 0")
	}

	if err := json.Unmarshal(reqData, &userData); err != nil {
		return core_errors.ServerError("error to decode json")
	}

	return nil
}
