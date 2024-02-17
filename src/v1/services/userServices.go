package services

import (
	"fmt"
	"net/http"
	"strconv"

	enc "github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/internal/passwd"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
)

func GetUser(db *pg.DB, id string) (enc.Response) {
  _int_id, e := strconv.Atoi(id)
  if e != nil {
    return err.BadRequestResponse("Provided Id is malformed")
  }
	user := repo.User{Id: _int_id}
	er := db.Model(&user).WherePK().Select()
	if er != nil {
    return err.NotFoundErrorResponse()
	}
  return enc.Response{Code: http.StatusOK, Payload: user}
}

func GetUsers(db *pg.DB, query string) (enc.Response) {
	var users []repo.User
  var er error
  if query != "" {
    // custom query
    // the reasoning being a wrong formatting on the end of go-pg when using query string
    er = db.Model(&users).Where(fmt.Sprintf("name ilike '%%%v%%'", query)).Select()
  } else {
    er = db.Model(&users).Select()
  }
	if er != nil {
    // we shouldn't get any errors here unless the table users doesn't exist
    return err.NotFoundErrorResponse(er.Error())
	}
  return enc.Response{Code: http.StatusOK, Payload: users}
}

func GetUsersCount(db *pg.DB) (enc.Response) {
	count, er := db.Model(&repo.User{}).Count()
	if er != nil {
    // we shouldn't get any errors here unless the table users doesn't exist
    return err.NotFoundErrorResponse()
	}
  return enc.Response{Code: http.StatusOK, Payload: count}
}

func DeleteUser(db *pg.DB, id string) (enc.Response) {
  _int_id, e := strconv.Atoi(id)
  if e != nil {
    return err.BadRequestResponse("Provided Id is malformed")
  }
	u := repo.User{Id: _int_id}
	_, er := db.Model(&u).WherePK().Delete()
	if er != nil {
    return err.NotFoundErrorResponse()
	}
  return enc.Response{Code: http.StatusOK}
}

func UpdateUser(db *pg.DB, u repo.User) (enc.Response) {
  // only name or phone can be changed
  _, er := db.Model(&u).Column("name", "phone").WherePK().UpdateNotZero()
	if er != nil {
    // the user wasn't found
    return err.NotFoundErrorResponse(er.Error())
	}
  return enc.Response{Code: http.StatusOK}
}

func AddUser(db *pg.DB, u repo.User) (enc.Response) {
  // NOTE: the following command allows us to get atomicity with transactions from Postgresql
  tx, er := db.Begin()
  defer tx.Close()
  x, er := tx.Model(&u).Insert()
	if er != nil {
    _ = tx.Rollback()
    return err.BadRequestResponse(er.Error())
	}
  if u.Role == "student" {
    s := repo.Student{Id: u.Id, Class: u.Class}
    _, er := tx.Model(&s).Insert()
    if er != nil {
      // this should never fail unless students table doesn't exist
      _ = tx.Rollback()
      return err.InternalServerErrorResponse(er.Error())
    }
  }
  // the default password is their respective phone number
  password := passwd.GetHash(u.Phone)
  c := repo.Credential{Id: strconv.Itoa(u.Id), Password: password, Role: u.Role}
  _, er = tx.Model(&c).Insert()
  if er != nil {
    _ = tx.Rollback()
    // this should never fail unless credentials table doesn't exist
    return err.InternalServerErrorResponse(er.Error())
  }
  if er := tx.Commit(); er != nil {
    panic(er.Error())
  }
  return enc.Response{Code: http.StatusCreated, Payload: x}
}
