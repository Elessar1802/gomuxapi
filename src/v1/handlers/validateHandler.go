package handlers

import (
	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/token"
  "net/http"
)
func (h Handlers) ValidateHandler(w http.ResponseWriter, r *http.Request) {
  claims := token.GetClaims(token.GetToken(r))
  // if the request has reached thus far we can say we have a valid token
  encoder.NewEncoder(w).Encode(encoder.Response{Code: http.StatusOK, Payload: claims["id"]})
}
