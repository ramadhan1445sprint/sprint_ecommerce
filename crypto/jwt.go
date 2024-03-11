package crypto

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ramadhan1445sprint/sprint_ecommerce/config"
)

func GenerateToken(username, name string) (string, error) {
	secret := config.GetString("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"name":     name,
		"exp":      jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
	})

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}
