package jwt

import (
	"log"
	"net/http"
	"time"

	"github.com/byte3/bookclub/backend/config"
	"github.com/go-chi/jwtauth"
)

type JWT struct {
	tokenClaim string
	tokenAuth  *jwtauth.JWTAuth
}

func (JWT) New() *JWT {
	jwt := &JWT{
		tokenClaim: "user_id",
		tokenAuth:  jwtauth.New("HS256", []byte(config.GetConfig().TokenSecretKey), nil),
	}
	log.Println("DEBUG JWT:", jwt.Encode("1"))
	return jwt
}

func (jwt *JWT) Encode(id string) string {
	claims := jwtauth.Claims{}.
		Set(jwt.tokenClaim, id).
		SetExpiryIn(30 * time.Second).
		SetIssuedNow()
	_, tokenString, _ := jwt.tokenAuth.Encode(claims)
	return tokenString
}

func (jwt *JWT) Verifier() func(http.Handler) http.Handler {
	return jwtauth.Verifier(jwt.tokenAuth)
}
