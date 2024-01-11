package router

import (
	"github.com/Elessar1802/api/src/v1/students"
	"github.com/gorilla/mux"
)

const API_PREFIX = "/api/v1"

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// using subrouters optimises request matching
	s := r.PathPrefix(API_PREFIX).Subrouter()

	// register the different routes
	students.InitStudentsSubrouter(s)

	return r
}
