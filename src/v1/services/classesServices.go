package services

import (
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func GetClasses(db *pg.DB) ([]string, *err.Error) {
	var classes []string
	er := db.Model(&repo.Student{}).Column("class").Distinct().Select(&classes)
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return classes, nil
}

func GetClass(db *pg.DB, name string) ([]repo.User, *err.Error) {
  users := []repo.User{}
  er := db.Model().Table("users").
    ColumnExpr("users.id, users.name, users.phone").
    Join("JOIN students ON students.id = users.id and students.class = ?", name).
    Select(&users)
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return users, nil
}
