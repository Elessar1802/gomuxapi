package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/services"
	"github.com/gorilla/mux"
)

func (h Handlers) ClassesHandler(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error
	method := r.Method

	switch method {
	case http.MethodGet:
		// return all the registered classes and their respective ids
		res, er = services.GetClasses(h.DB)

	case http.MethodPost:
		class := repo.Class{}
		json.NewDecoder(r.Body).Decode(&class)
		res, er = services.AddClass(h.DB, class)

	default:
		res = nil
		er = &err.Error{
			Code:    400,
			Message: "Bad request!",
		}
	}

	encoder.NewEncoder(w).Encode(res, er)
}

func (h Handlers) ClassesHandlerId(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error

	params := mux.Vars(r)
	id := params["id"]
	method := r.Method

	switch method {
	case http.MethodGet:
		res, er = services.GetClass(h.DB, id)

	default:
		res = nil
		er = &err.Error{
			Code:    400,
			Message: "Bad request!",
		}
	}

	encoder.NewEncoder(w).Encode(res, er)
}
