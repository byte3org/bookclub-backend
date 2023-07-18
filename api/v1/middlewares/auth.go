package middlewares

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func ensureAdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("token sign method is invalid"))
					if err != nil {
						return nil, err
					}
				}
				return "", nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("error with jwt"))
				if err != nil {
					return
				}
			}

			ctx := context.WithValue(r.Context(), "authenticated", "admin")

			if token.Valid {
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		}
	})
}

func ensureAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {

		}
	})
}
