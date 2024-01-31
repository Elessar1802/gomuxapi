package services

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func GetClasses(db *pg.DB) (encoder.Response) {
	var classes []string
	er := db.Model(&repo.Student{}).Column("class").Distinct().Select(&classes)
	if er != nil {
		return err.NotFoundErrorResponse()
	}
  return encoder.Response{Code: http.StatusOK, Payload: classes}
}

func GetClass(db *pg.DB, name string) (encoder.Response) {
  users := []repo.User{}
  er := db.Model().Table("users").
    ColumnExpr("users.id, users.name, users.phone").
    Join("JOIN students ON students.id = users.id and students.class = ?", name).
    Select(&users)
	if er != nil || len(users) == 0 {
		return err.NotFoundErrorResponse("Class `" + name + "` not found")
	}
  return encoder.Response{Code: http.StatusOK, Payload: users}
}
