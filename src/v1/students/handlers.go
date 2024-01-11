package students

import (
	"encoding/json"
	"github.com/Elessar1802/api/src/v1/internal/encoder"
	"github.com/Elessar1802/api/src/v1/internal/err"
	"github.com/gorilla/mux"
	"net/http"
)

var students = []student{
	student{"1802", "Amrit Borah", "443 Purbalok, Kolkata 781036", 10, "A"},
	student{"1802", "Amrit Borah", "443 Purbalok, Kolkata 781036", 10, "A"},
	student{"1802", "Amrit Borah", "443 Purbalok, Kolkata 781036", 10, "A"},
	student{"1802", "Amrit Borah", "443 Purbalok, Kolkata 781036", 10, "A"},
}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder.NewEncoder(w).Encode(students, nil)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := params["id"]

	// find the right student
	for _, s := range students {
		if s.Id != id {
			continue
		}
		w.WriteHeader(http.StatusOK)
		encoder.NewEncoder(w).Encode(s, nil)
		return
	}

	// TODO: use the right code
	w.WriteHeader(http.StatusBadRequest)
	encoder.NewEncoder(w).
		Encode(nil,
			&err.Error{
				Code:    http.StatusBadRequest,
				Message: "Student not found!",
			})
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	// TODO: I need to use a middleware
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	s := student{}
	json.NewDecoder(r.Body).Decode(&s)
	s.Id = "1920"
	students = append(students, s)
	encoder.NewEncoder(w).Encode(s, nil)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	id := params["id"]
	var std *student
	for idx, _ := range students {
		s := &students[idx]
		if s.Id != id {
			continue
		}
		std = s
		break
	}
	if std == nil {
		return
	}
	json.NewDecoder(r.Body).Decode(&std)
	encoder.NewEncoder(w).Encode(std, nil)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := params["id"]

	// find the right student
	del := -1
	for idx, s := range students {
		if s.Id != id {
			continue
		}
		del = idx
		w.WriteHeader(http.StatusOK)
		encoder.NewEncoder(w).Encode(nil, nil)
		break
	}

	if del != -1 {
		students[del] = students[len(students)-1]
		students = students[:len(students)-1]
		return
	}

	// TODO: use the right code
	w.WriteHeader(http.StatusBadRequest)
	encoder.NewEncoder(w).
		Encode(nil,
			&err.Error{
				Code:    http.StatusBadRequest,
				Message: "Student not found!",
			})
}

func getAttendance(w http.ResponseWriter, r *http.Request) {

}

func punchIn(w http.ResponseWriter, r *http.Request) {

}

func punchOut(w http.ResponseWriter, r *http.Request) {

}
