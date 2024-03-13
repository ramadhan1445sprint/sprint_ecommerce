package entity

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id       string `db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JWTPayload struct {
	Id       string
	Username string
	Name     string
}

type JWTClaims struct {
	Id       string
	Username string
	Name     string
	jwt.RegisteredClaims
}
