package helpers

import (
	"time"

	"github.com/byte3/bookclub/backend/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GenerateAuthToken(username string, email string) (string, error) {
	secret_key := config.GetConfig().TokenSecretKey

	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["email"] = email
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(240 * time.Hour)

	tokenString, err := token.SignedString(secret_key)

	return tokenString, err
}

func ExtractClaims(token string) {

}

func VerifyAuthToken(token string) {

}
