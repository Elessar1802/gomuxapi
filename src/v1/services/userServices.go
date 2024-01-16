package services

import (
	"reflect"
  "strings"

	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func GetUser(db *pg.DB, id string) (interface{}, *err.Error) {
	user := repo.User{Id: id}
	er := db.Model(&user).WherePK().Select()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return user, nil
}

func GetUsers(db *pg.DB) ([]repo.User, *err.Error) {
	var users []repo.User
	er := db.Model(&users).Select()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return users, nil
}

func DeleteUser(db *pg.DB, id string) (interface{}, *err.Error) {
	u := repo.User{Id: id}
	res, er := db.Model(&u).WherePK().Delete()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return res, nil
}

func UpdateUser(db *pg.DB, u repo.User) (interface{}, *err.Error) {
  mp := map[string]interface{}{}
  v := reflect.ValueOf(u)
  s := v.Type()
  empty := repo.User{}
  ev := reflect.ValueOf(empty)
  for i := 0; i < v.NumField(); i++ {
    if v.Field(i).Interface() == ev.Field(i).Interface() {
      continue
    }
    mp[strings.ToLower(s.Field(i).Name)] = v.Field(i).Interface()
  }
  res, er := db.Model(&mp).TableExpr("users").Where("id = ?", u.Id).Update()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return res, nil
}

func AddUser(db *pg.DB, u repo.User) (interface{}, *err.Error) {
	_, er := db.Model(&u).Insert()
	if er != nil {
		return nil, &err.Error{Code: 404, Message: er.Error()}
	}
	return u, nil
}
