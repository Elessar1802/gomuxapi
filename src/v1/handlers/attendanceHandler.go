package handlers

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/services"
  "github.com/gorilla/mux"
)

func (h Handlers) AttendanceClassHandlerId(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error
	method := r.Method
  params := mux.Vars(r)
  id := params["id"]
  from := params["start_date"]
  to := params["end_date"]

	switch method {
	case http.MethodGet:
		// return all the registered classes and their respective ids
    res, er = services.AttendanceClassId(h.DB, id, from, to)

	default:
		res = nil
		er = &err.Error{
			Code:    400,
			Message: "Bad request!",
		}
	}

	encoder.NewEncoder(w).Encode(res, er)
}

func (h Handlers) AttendanceUserHandlerId(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error
	method := r.Method
  params := mux.Vars(r)
  id := params["id"]
  from := params["start_date"]
  to := params["end_date"]

	switch method {
	case http.MethodGet:
		// return all the registered classes and their respective ids
    res, er = services.AttendanceUserId(h.DB, id, from, to)

  case http.MethodPost:
    res, er = services.PunchIn(h.DB, id)

  case http.MethodPut:
    res, er = services.PunchOut(h.DB, id)

	default:
		res = nil
		er = &err.Error{
			Code:    400,
			Message: "Bad request!",
		}
	}

	encoder.NewEncoder(w).Encode(res, er)
}
