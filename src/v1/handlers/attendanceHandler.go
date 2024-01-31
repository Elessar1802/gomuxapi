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
	var res encoder.Response
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
			res = err.UnauthorizedAccessResponse()
			break
		}
		// return all the registered classes and their respective ids
		res = services.AttendanceClassId(h.DB, id, from, to)

	default:
		res = err.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
}

func (h Handlers) AttendanceUserHandlerId(w http.ResponseWriter, r *http.Request) {
	var res encoder.Response
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
    // we are going to fetch the details of user whose attendance is being requested 
    // after fetching it we are going to run it through a check whether the user making the request
    // is authorized to do so
    r := services.GetUser(h.DB, id)
    if r.Error != nil {
      res = r
      break
    }
    user := r.Payload
    u, ok := user.(repo.User)
		if !ok {
			res = err.InternalServerErrorResponse()
			break
		}
		_id, _ := claims["id"].(string)
		_role, _ := claims["role"].(string)
		if authorized := isAuthorizedToViewAttendance(_id, _role, u); !authorized {
			res = err.UnauthorizedAccessResponse()
      break
		}
    res = services.AttendanceUserId(h.DB, id, from, to)

	case http.MethodPost:
		// only teachers and students can register their attendance
		if (claims["role"] != "teacher" && claims["role"] != "student") || claims["id"] != id {
      // only the user can punch in/out their own attendance
			res = err.UnauthorizedAccessResponse()
			break
		}
		res = services.PunchIn(h.DB, id)

	case http.MethodPut:
		// only teachers and students can register their attendance
		if (claims["role"] != "teacher" && claims["role"] != "student") || claims["id"] != id {
      // only the user can punch in/out their own attendance
			res = err.UnauthorizedAccessResponse()
			break
		}
		res = services.PunchOut(h.DB, id)

	default:
		res = err.MethodNotAllowedErrorResponse()
	}

	encoder.NewEncoder(w).Encode(res)
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
