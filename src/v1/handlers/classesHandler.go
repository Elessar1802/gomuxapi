package handlers

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/services"
	"github.com/gorilla/mux"
)

func (h Handlers) ClassesHandler(w http.ResponseWriter, r *http.Request) {
	var res encoder.Response
	method := r.Method

	switch method {
	case http.MethodGet:
		// return all the registered classes and their respective ids
		res = services.GetClasses(h.DB)

	default:
    res = err.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}

func (h Handlers) ClassesHandlerId(w http.ResponseWriter, r *http.Request) {
	var res encoder.Response

	params := mux.Vars(r)
	id := params["id"]
	method := r.Method

	switch method {
	case http.MethodGet:
		res = services.GetClass(h.DB, id)

	default:
    res = err.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}
