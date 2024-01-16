package services

import (
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func GetClasses(db *pg.DB) ([]repo.Class, *err.Error) {
	var classes []repo.Class
	er := db.Model(&classes).Select()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return classes, nil
}

/*
 * func will return (@repo.Class, nil) if no error
 * otherwise (nil, err)
 */
func AddClass(db *pg.DB, c repo.Class) (interface{}, *err.Error) {
	_, er := db.Model(&c).Insert()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return c, nil
}

func GetClass(db *pg.DB, name string) (interface{}, *err.Error) {
	cl := repo.Class{Name: name}
	er := db.Model(&cl).WherePK().Select()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return cl, nil
}
