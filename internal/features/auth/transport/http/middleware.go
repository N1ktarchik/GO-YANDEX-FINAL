package http

import (
	"net/http"
)

func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ValidPass := h.authService.GetPath()

		if len(ValidPass) > 0 {
			var jwt string

			cookie, err := r.Cookie("token")
			if err == nil {
				jwt = cookie.Value
			}

			claims, err := h.authService.ValidateJWT(jwt)

			if err != nil {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

			if !h.authService.Compare(claims.HashPass) {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}
