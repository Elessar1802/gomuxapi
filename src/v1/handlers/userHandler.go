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
	var res interface{}
	var err *er.Error
	token := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	// we are parsing unverified as the middleware verifies the token
  jwt.NewParser().ParseUnverified(token, claims)

	switch method {
	case http.MethodGet:
		res, err = services.GetUsers(h.DB)
	case http.MethodPost:
    if claims["role"] != "principal" {
      res, err = er.UnauthorizedAccessResponse()
      break
    }
		user := repo.User{}
		json.NewDecoder(r.Body).Decode(&user)
    if user.Role == "student" && user.Class == "" {
      res, err = er.BadRequestResponse()
      break
    }
    user.Id = uuid.NewString()
		res, err = services.AddUser(h.DB, user)
	}

	encoder.NewEncoder(w).Encode(res, err)
}

func (h Handlers) UsersHandlerId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	method := r.Method
	var res interface{}
	var err *er.Error
	switch method {
	case http.MethodGet:
		res, err = services.GetUser(h.DB, id)
	case http.MethodPut:
		update := repo.User{}
		json.NewDecoder(r.Body).Decode(&update)
		update.Id = id
		res, err = services.UpdateUser(h.DB, update)
	case http.MethodDelete:
		res, err = services.DeleteUser(h.DB, id)
	default:
		res, err = er.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res, err)
}
