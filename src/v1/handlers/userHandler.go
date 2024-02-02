package handlers

import (
	"encoding/json"
	"net/http"
  "strconv"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	er "github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/services"
	"github.com/gorilla/mux"
)

func (h Handlers) UsersHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	var res encoder.Response
	
	switch method {
	case http.MethodGet:
		res = services.GetUsers(h.DB)
	case http.MethodPost:
		user := repo.User{}
		json.NewDecoder(r.Body).Decode(&user)
    if user.Role == "student" && user.Class == "" {
      res = er.BadRequestResponse()
      break
    }
    // create a uuid
		res = services.AddUser(h.DB, user)
  default:
		res = er.MethodNotAllowedErrorResponse()
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
    _int_id, e := strconv.Atoi(id)
    if e != nil {
      res = er.BadRequestResponse("Provided id is malformed")
    }
		update.Id = _int_id
		res = services.UpdateUser(h.DB, update)
	case http.MethodDelete:
		res = services.DeleteUser(h.DB, id)
	default:
		res = er.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}
