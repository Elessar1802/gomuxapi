package services

import (
	"net/http"
	"os"

	enc "github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/internal/passwd"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/go-pg/pg/v10"
	jwt "github.com/golang-jwt/jwt/v5"
)

func generateUserToken(u repo.Credential) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   u.Id,
		"role": u.Role,
	})
	secret := os.Getenv("JWT_SECRET")
	tokenString, e := token.SignedString([]byte(secret))
	if e != nil {
		return "", e
	}
	return tokenString, nil
}

func SetToken(db *pg.DB, c repo.Credential) enc.Response {
	var token string
	var er interface{}
  /* to ensure password safety we store the base64 encoded
   * hashes of the passwords and not the passwords themselves */
	password := passwd.GetHash(c.Password)

	// fill the fields of the struct with the db data and check if the passwords match
	e := db.Model(&c).Where("id = ?id").Select()
	if e != nil || c.Password != password {
		return err.BadRequestResponse("Either username or password is wrong")
	}
	if c.Token == "" {
		token, er = generateUserToken(c)
		if er != nil {
			return err.InternalServerErrorResponse("Unable to create a jwt token")
		}
		c.Token = token
		_, er = db.Model(&c).Column("token").Where("id = ?id").Update()
		if er != nil {
			return err.InternalServerErrorResponse()
		}
	}

  cookie := http.Cookie{Name: "app.token", Value: c.Token, Path: "/api/v1", MaxAge: 5000000}
	return enc.Response{Code: http.StatusOK, Payload: c.Token, Cookie: &cookie}
}

func DeleteToken(db *pg.DB) enc.Response {
  cookie := http.Cookie{Name: "app.token", Value: "", MaxAge: -1};
	return enc.Response{Code: http.StatusOK, Cookie: &cookie}
}

func UpdatePassword(db *pg.DB, c repo.Credential) enc.Response {
	// only password can be changed
  /* to ensure password safety we store the base64 encoded
   * hashes of the passwords and not the passwords themselves */
	c.Password = passwd.GetHash(c.Password)
	_, e := db.Model(&c).Column("password").Where("id = ?id").UpdateNotZero()
	if e != nil {
		return err.BadRequestResponse("Malformed credential object")
	}
	return enc.Response{Code: http.StatusOK}
}
