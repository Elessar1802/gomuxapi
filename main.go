package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/router"
	"github.com/Elessar1802/api/src/v1/services"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
)

const PORT = ":8000"
const API_PREFIX = "/api/v1"

func main() {
  // the kube service needs to be setup before the pod is created. Otherwise these env vars won't be available
  // if we have coredns setup we can directly specify the service.metadata.name for hostname and then separately provide the port
  addr := fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
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
  // there is no orm mapping for alter table hence need to use Exec
  db.Model().Exec("ALTER TABLE students ADD CONSTRAINT FK_UserId FOREIGN KEY(id) REFERENCES users(id)");
  db.Model().Exec("ALTER TABLE attendance ADD CONSTRAINT FK_UserId FOREIGN KEY(id) REFERENCES users(id)");

  // initialize the principal user
  name := os.Getenv("PRINCIPAL_NAME")
  phone := os.Getenv("PRINCIPAL_PHONE")
  if name == "" || phone == "" {
    panic("Missing or malformed principal details as ENV VARS")
  }
  principal := repo.User{Id: 1, Name: name, Phone: phone, Role: "principal"}
  fmt.Println(services.AddUser(db, principal).Payload) // we are returning the token
}
