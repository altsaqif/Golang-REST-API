package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte(GoDotEnvVariable("SECRET_KEY"))

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
