package handlers

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/services"
	"github.com/gorilla/mux"
)

func (h Handlers) AttendanceClassHandlerId(w http.ResponseWriter, r *http.Request) {
	var res encoder.Response
	method := r.Method
	params := mux.Vars(r)
	id := params["id"]
	from := params["start_date"]
	to := params["end_date"]

	switch method {
	case http.MethodGet:
		// return all the registered classes and their respective ids
		res = services.AttendanceClassId(h.DB, id, from, to)

	default:
		res = err.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}

func (h Handlers) AttendanceUserHandlerId(w http.ResponseWriter, r *http.Request) {
	var res encoder.Response
	method := r.Method
	params := mux.Vars(r)
	id := params["id"]
	from := params["start_date"]
	to := params["end_date"]

	switch method {
	case http.MethodGet:
    res = services.AttendanceUserId(h.DB, id, from, to)

	case http.MethodPost:
    // no check here since we already handle it before this handler is called
		res = services.PunchIn(h.DB, id)

	case http.MethodPut:
    // no check here since we already handle it before this handler is called
		res = services.PunchOut(h.DB, id)

	default:
		res = err.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}
