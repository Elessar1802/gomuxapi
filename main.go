package main

import (
	"log"
	"net/http"

	"github.com/Elessar1802/api/src/v1/router"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
)

const PORT = ":8000"
const API_PREFIX = "/api/v1"

func main() {
  db := pg.Connect(&pg.Options{
		User: "amloch",
    Database: "app",
	})
	defer db.Close()
	r := mux.NewRouter()

	// using subrouters optimises request matching
	s := r.PathPrefix(API_PREFIX).Subrouter()

	// register the different routes
	router.InitRouter(s, db)

	// start the server
	log.Fatal(http.ListenAndServe(PORT, r))
}
