package main

import (
	"log"
	"net/http"
	"os"

	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/router"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
)

const PORT = ":8000"
const API_PREFIX = "/api/v1"

func main() {
  addr := os.Getenv("DB_ADDR")
  db := pg.Connect(&pg.Options{
		User: "postgres",
    Password: "odin",
    Database: "postgres",
    Addr: addr,
	})
  createTables(db)
	defer db.Close()
	r := mux.NewRouter()

	// using subrouters optimises request matching
	s := r.PathPrefix(API_PREFIX).Subrouter()

	// register the different routes
	router.InitRouter(s, db)

	// start the server
	log.Fatal(http.ListenAndServe(PORT, r))
}

func createTables (db *pg.DB) {
  opts := &orm.CreateTableOptions{
    IfNotExists: true,
  }
  db.Model(&repo.User{}).CreateTable(opts)
  db.Model(&repo.Attendance{}).CreateTable(opts)
  db.Model(&repo.Student{}).CreateTable(opts)
}
