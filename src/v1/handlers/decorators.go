package handlers

import (
	"net/http"
  "strconv"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/Elessar1802/api/src/v1/services"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/gorilla/mux"
  "github.com/Elessar1802/api/src/v1/internal/token"
)

/*
 * No authorization is handled inside the handlers themselves
 * Instead we use a decorator function pattern to do authorization
 * See: src/v1/router/router.go
 */

func (h Handlers) OnlyPrincipal(fn http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    claims := token.GetClaims(token.GetToken(r))

    if claims["role"] != "principal" {
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
    }

    fn.ServeHTTP(w, r)
  })
}

func (h Handlers) OnlyMatchingID(fn http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    claims := token.GetClaims(token.GetToken(r))

    if id != claims["id"] {
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
    }

    fn.ServeHTTP(w, r)
  })
}

func (h Handlers) NotPrincipal(fn http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    claims := token.GetClaims(token.GetToken(r))

    if "principal" == claims["role"] {
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
    }

    fn.ServeHTTP(w, r)
  })
}

func (h Handlers) OnlyTeachers(fn http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    claims := token.GetClaims(token.GetToken(r))

    if claims["role"] != "teacher" {
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
    }

    fn.ServeHTTP(w, r)
  })
}

func (h Handlers) UserAttendanceChecks(fn http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    claims := token.GetClaims(token.GetToken(r))

    // we are going to fetch the details of user whose attendance is being requested 
    // after fetching it we are going to run it through a check whether the user making the request
    // is authorized to do so
    u := services.GetUser(h.DB, id)
    if u.Error != nil {
      encoder.NewEncoder(w).Encode(err.BadRequestResponse(u.Error.(*err.Error).Message))
      return
    }
    user := u.Payload.(repo.User)
    _int_id, _ := strconv.Atoi(claims["id"].(string))
    if authorized := isAuthorizedToViewAttendance(_int_id, claims["role"].(string), user); !authorized {
      encoder.NewEncoder(w).Encode(err.UnauthorizedAccessResponse())
      return
		}

    fn.ServeHTTP(w, r)
  })
}

func isAuthorizedToViewAttendance(id int, role string, user repo.User) bool {
	if (role == "principal" && user.Role != "teacher") || // a principal can only view a teacher's attendance
		 (role == "teacher" && user.Role == "teacher" && id != user.Id) || // a teacher can view student attendance and as well their own attendance
		 (user.Role == "principal") || // a principal doesn't have attendance records
		 (role == "student" && id != user.Id) { // a student can only view their own attendance
		return false
	}
	return true
}
