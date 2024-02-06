package token

import (
	"net/http"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GetToken(r *http.Request) string {
  cookie, err := r.Cookie("app.token")
  if err == nil {
    return cookie.Value
  }
  token := r.Header.Get("Authorization")
  return token
}

func GetClaims(token string) jwt.MapClaims {
  if token == "" {
    return nil
  }
  claims := jwt.MapClaims{}
  jwt.NewParser().ParseUnverified(token, claims)
  return claims
}
