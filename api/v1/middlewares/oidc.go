package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/byte3/bookclub/backend/config"
	"github.com/byte3/bookclub/backend/helpers"
)

func ExtractOidcToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		oidcToken := authHeader[1]
		ctx := context.WithValue(r.Context(), "oidc", oidcToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ExtractUserIdfromOidc(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // extract oidc token from request context
        // do the client things
        token := r.Context().Value("token").(string)
        
        client := http.Client{}
        req, err := http.NewRequest("GET", config.GetConfig().BookclubAuth, nil)    
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        req.Header = http.Header{
            "Authorization" : {"Bearer " + token},
        }
        res, err := client.Do(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer res.Body.Close()
        jsonRes, err := helpers.ParseJsonBody(res.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        userId := jsonRes["sub"].(int)
        ctx := context.WithValue(r.Context(), "user_id", userId)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
