package router

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/handlers"
	"github.com/Elessar1802/api/src/v1/internal/middleware"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

type Route struct {
	path    string
	handler http.HandlerFunc
	methods []string
	queries []string // format - key[0], value[0], key[1], value[1] ...
}

func getRoutes(db *pg.DB) []Route {
  h := &handlers.Handlers{DB: db}
  return []Route{
    {"/users", h.UsersHandler, []string{"GET", "POST"}, nil},
    {"/users/{id}", h.UsersHandlerId, []string{"GET", "PUT", "DELETE"}, nil},
    // TODO: ask if this is a valid way of using query string
    {"/attendance/class/{id}", h.AttendanceClassHandlerId, []string{"GET"}, []string{"start_date", "{start_date}", "end_date", "{end_date}"}},
    {"/attendance/user/{id}", h.AttendanceUserHandlerId, []string{"GET"}, []string{"start_date", "{start_date}", "end_date", "{end_date}"}},
    {"/attendance/user/{id}", h.AttendanceUserHandlerId, []string{"POST", "PUT"}, nil},
    {"/classes", h.ClassesHandler, []string{"GET", "POST"}, nil},
    {"/classes/{id}", h.ClassesHandlerId, []string{"GET"}, nil},
  }
}

func InitRouter(router *mux.Router, db *pg.DB) {
	for _, r := range getRoutes(db) {
		router.
			Path(r.path).
			HandlerFunc(r.handler).
			Methods(r.methods...).
			Queries(r.queries...)
	}
	router.Use(middleware.Jsonmw)
	// authentication
	router.Use(middleware.Auth)
}
