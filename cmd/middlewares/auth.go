package middlewares

import (
	"net/http"

	"github.com/altsaqif/go-restapi-mux/cmd/helpers"
	"github.com/golang-jwt/jwt/v4"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Nangkap Token Dari Cookie
		accessToken, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helpers.Response(w, http.StatusUnauthorized, response, nil)
				return
			}
		}

		// Mengambil Token Value
		tokenString := accessToken.Value

		claims := &helpers.JWTClaim{}

		// Parsing Token JWT
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return helpers.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {

			case jwt.ValidationErrorSignatureInvalid:
				// Token Invalid
				response := map[string]string{"message": "Unauthorized"}
				helpers.Response(w, http.StatusUnauthorized, response, nil)
				return

			case jwt.ValidationErrorExpired:
				// Token Expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helpers.Response(w, http.StatusUnauthorized, response, nil)
				return

			default:
				response := map[string]string{"message": "Unauthorized"}
				helpers.Response(w, http.StatusUnauthorized, response, nil)
				return
			}
		}

		if !token.Valid {
			response := map[string]string{"message": "Unauthorized"}
			helpers.Response(w, http.StatusUnauthorized, response, nil)
			return
		}

		// ctx := context.WithValue(r.Context(), "userinfo", tokenString)
		// ctx := r.Context()
		// ctx = context.WithValue(ctx, "userID", claims.UserID)
		// next.ServeHTTP(w, r.WithContext(ctx))
		next.ServeHTTP(w, r)
	})
}
