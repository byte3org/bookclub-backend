package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RequestCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Context middleware called")
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
