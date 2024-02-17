package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	er "github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/services"
)

func (h Handlers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	var res encoder.Response

	switch method {
  // Why we should use POST for sending Credential to server for login?
  // https://stackoverflow.com/questions/5868786/what-method-should-i-use-for-a-login-authentication-request
	case http.MethodPost:
		c := repo.Credential{}
		e := json.NewDecoder(r.Body).Decode(&c)
		if e != nil {
			res = er.BadRequestResponse(e.Error())
			break
		}
		res = services.SetToken(h.DB, c)
	case http.MethodPut:
		c := repo.Credential{}
		e := json.NewDecoder(r.Body).Decode(&c)
		if e != nil {
			res = er.BadRequestResponse(e.Error())
			break
		}
		res = services.UpdatePassword(h.DB, c)
  case http.MethodDelete:
    res = services.DeleteToken(h.DB)
	default:
		res = er.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}
