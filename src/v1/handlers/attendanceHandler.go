package handlers

import (
	"net/http"

	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	repo "github.com/Elessar1802/api/src/v1/repository"
	"github.com/Elessar1802/api/src/v1/services"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func (h Handlers) AttendanceClassHandlerId(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error
	method := r.Method
	params := mux.Vars(r)
	id := params["id"]
	from := params["start_date"]
	to := params["end_date"]
	token := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	// we are parsing unverified as the middleware verifies the token
  jwt.NewParser().ParseUnverified(token, claims)

	switch method {
	case http.MethodGet:
		// only a teacher can check the attendance of
		if claims["role"] != "teacher" {
			res, er = err.UnauthorizedAccessResponse()
			break
		}
		// return all the registered classes and their respective ids
		res, er = services.AttendanceClassId(h.DB, id, from, to)

	default:
		res, er = err.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res, er)
}

func (h Handlers) AttendanceUserHandlerId(w http.ResponseWriter, r *http.Request) {
	var res interface{}
	var er *err.Error
	method := r.Method
	params := mux.Vars(r)
	id := params["id"]
	from := params["start_date"]
	to := params["end_date"]
	token := r.Header.Get("Authorization")
	claims := jwt.MapClaims{}
	jwt.NewParser().ParseUnverified(token, claims)

	switch method {
	case http.MethodGet:
    // the result will be of type []repo.Attendance
    user, x := services.GetUser(h.DB, id)
    if x != nil {
      res, er = err.BadRequestResponse()
      break
    }
    u, ok := user.(repo.User)
		if !ok {
			res, er = err.BadRequestResponse()
			break
		}
		i, _ := claims["id"].(string)
		r, _ := claims["role"].(string)
		if authorized := isAuthorizedToViewAttendance(i, r, u); !authorized {
			res, er = err.UnauthorizedAccessResponse()
      break
		}
		res, er = services.AttendanceUserId(h.DB, id, from, to)

	case http.MethodPost:
		// only teachers and students can register their attendance
		if (claims["role"] != "teacher" && claims["role"] != "student") || claims["id"] != id {
      // only the user can punch in/out their own attendance
			res, er = err.UnauthorizedAccessResponse()
			break
		}
		res, er = services.PunchIn(h.DB, id)

	case http.MethodPut:
		// only teachers and students can register their attendance
		if (claims["role"] != "teacher" && claims["role"] != "student") || claims["id"] != id {
      // only the user can punch in/out their own attendance
			res, er = err.UnauthorizedAccessResponse()
			break
		}
		res, er = services.PunchOut(h.DB, id)

	default:
		res, er = err.BadRequestResponse()
	}

	encoder.NewEncoder(w).Encode(res, er)
}

func isAuthorizedToViewAttendance(id string, role string, user repo.User) bool {
	if (role == "principal" && user.Role != "teacher") || // a principal can only view a teacher's attendance
		(role == "teacher" && user.Role == "teacher" && id != user.Id) || // a teacher can view student attendance and as well their own attendance
		(user.Role == "principal") || // a principal doesn't have attendance records
		(role == "student" && id != user.Id) { // a student can only view their own attendance
		return false
	}
	return true
}
