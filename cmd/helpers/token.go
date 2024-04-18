package helpers

import (
	"time"

	"github.com/altsaqif/go-restapi-mux/cmd/configs"
	"github.com/altsaqif/go-restapi-mux/cmd/models"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte(configs.GoDotEnvVariable("SECRET_KEY"))

type JWTClaim struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateToken(user *models.Users) (string, error) {
	claims := JWTClaim{
		int(user.Model.ID),
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(JWT_KEY)
	return ss, err
}
