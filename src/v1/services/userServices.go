package services

import (
	"net/http"
	"os"
	"strconv"

	enc "github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
	jwt "github.com/golang-jwt/jwt/v5"
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

func GetUsers(db *pg.DB) (enc.Response) {
	var users []repo.User
	er := db.Model(&users).Select()
	if er != nil {
    // we shouldn't get any errors here unless the table users doesn't exist
    return err.NotFoundErrorResponse()
	}
  return enc.Response{Code: http.StatusOK, Payload: users}
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
  return enc.Response{Code: http.StatusNoContent}
}

func UpdateUser(db *pg.DB, u repo.User) (enc.Response) {
  _, er := db.Model(&u).WherePK().UpdateNotZero()
	if er != nil {
    // the user wasn't found
    return err.NotFoundErrorResponse()
	}
  return enc.Response{Code: http.StatusNoContent}
}

func generateUserToken(u repo.User) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id": u.Id,
    "role": u.Role,
  })
  secret := os.Getenv("JWT_SECRET")
  tokenString, e := token.SignedString([]byte(secret))
  if e != nil {
    return "", e
  }
  return tokenString, nil
}

func AddUser(db *pg.DB, u repo.User) (enc.Response) {
  token, e := generateUserToken(u)
  if e != nil {
    return err.InternalServerErrorResponse()
  }
	_, er := db.Model(&u).Insert()
	if er != nil {
    return err.BadRequestResponse()
	}
  if u.Role == "student" {
    s := repo.Student{Id: u.Id, Class: u.Class}
    _, er := db.Model(&s).Insert()
    if er != nil {
      // this should never fail unless students table doesn't exist
      return err.InternalServerErrorResponse()
    }
  }
  return enc.Response{Code: http.StatusCreated, Payload: token}
}
