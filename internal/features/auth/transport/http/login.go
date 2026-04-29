package http

import (
	core_errors "n1ktarchik/go-final/internal/core/errors"
	req "n1ktarchik/go-final/internal/core/transport/request"
	resp "n1ktarchik/go-final/internal/core/transport/response"
	"net/http"
)

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	authReq := &AuthDTO{}

	if err := req.DecodeJSON(r, authReq); err != nil {
		resp.RespondWithError(w, core_errors.BadRequest("error to decode JSON"))
	}

	token, err := h.authService.Login(authReq.Password)
	if err != nil {
		resp.RespondWithError(w, err)
	}

	resp.RespondWithJWT(w, http.StatusOK, token)
}
