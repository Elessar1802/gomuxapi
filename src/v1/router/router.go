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
    // If we are using Authorization header we need to configure OPTIONS for all endpoints
    // Otherwise we can get by without bothering for it if we are only bothered with simple requests
    // such as GET. For POST, PUT, DELETE (since they are non-simple need to configure OPTIONS method)
    // Options header is used for preflight requests from the fetchAPI
    {"/users", h.UsersHandler, []string{"GET", "POST", "OPTIONS"}, nil},
    {"/users/{id}", h.UsersHandlerId, []string{"GET", "PUT", "DELETE", "OPTIONS"}, nil},
    // TODO: ask if this is a valid way of using query string
    {"/attendance/class/{id}", h.AttendanceClassHandlerId, []string{"GET", "OPTIONS"}, []string{"start_date", "{start_date}", "end_date", "{end_date}"}},
    {"/attendance/user/{id}", h.AttendanceUserHandlerId, []string{"GET", "OPTIONS"}, []string{"start_date", "{start_date}", "end_date", "{end_date}"}},
    {"/attendance/user/{id}", h.AttendanceUserHandlerId, []string{"POST", "PUT", "OPTIONS"}, nil},
    {"/classes", h.ClassesHandler, []string{"GET", "POST", "OPTIONS"}, nil},
    {"/classes/{id}", h.ClassesHandlerId, []string{"GET", "OPTIONS"}, nil},
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
	router.Use(mux.CORSMethodMiddleware(router))
  // handling the options header in the middleware
  // We need only reply with a status 200 OK response
  // for the preflight to be successful
	router.Use(middleware.Corsmw)
	// authentication
	router.Use(middleware.Auth)
}
