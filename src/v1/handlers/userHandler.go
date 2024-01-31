package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	er "github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/services"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h Handlers) UsersHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	var res encoder.Response
	
	token := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	// we are parsing unverified as the middleware verifies the token
  jwt.NewParser().ParseUnverified(token, claims)

	switch method {
	case http.MethodGet:
		res = services.GetUsers(h.DB)
	case http.MethodPost:
    if claims["role"] != "principal" {
      res = er.UnauthorizedAccessResponse()
      break
    }
		user := repo.User{}
		json.NewDecoder(r.Body).Decode(&user)
    if user.Role == "student" && user.Class == "" {
      res = er.BadRequestResponse()
      break
    }
    user.Id = uuid.NewString()
		res = services.AddUser(h.DB, user)
	}

	encoder.NewEncoder(w).Encode(res)
}

func (h Handlers) UsersHandlerId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	method := r.Method
	var res encoder.Response
	
	switch method {
	case http.MethodGet:
		res = services.GetUser(h.DB, id)
	case http.MethodPut:
		update := repo.User{}
		json.NewDecoder(r.Body).Decode(&update)
		update.Id = id
		res = services.UpdateUser(h.DB, update)
	case http.MethodDelete:
		res = services.DeleteUser(h.DB, id)
	default:
		res = er.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}
