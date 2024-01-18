package middleware

import (
	"net/http"
	"os"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if not authenticated user then we cannot use
		token := r.Header.Get("Authorization")

    _, er := jwt.NewParser().Parse(token, func(t *jwt.Token) (interface{}, error) {
      secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})

    if er != nil {
      // we will stop propagtion if we can't verify the token
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
    }

		next.ServeHTTP(w, r)
	})
}
