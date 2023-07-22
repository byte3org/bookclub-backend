package jwt

import (
	"log"

	"github.com/byte3/bookclub/backend/config"
	"github.com/go-chi/jwtauth"
)

var (
	TokenAuth *jwtauth.JWTAuth
)

func Init() {
	TokenAuth = jwtauth.New("HS256", []byte(config.GetConfig().TokenSecretKey), nil)
}

func GenerateToken(id int, username string, email string) string {
	_, token_str := TokenAuth.Encode(map[string]interface{}{
		"user_id":  id,
		"username": username,
		"email":    email,
	})
	log.Println("[!] token generated :", token_str)
	return token_str
}
