package students

import (
	"github.com/gorilla/mux"
	"net/http"
)

type query struct {
	key   string
	value string
}

type route struct {
	path    string
	handler http.HandlerFunc
	method  string
	queries []query
}

func getRoutes() []route {
	return []route{
		route{"/", getAllStudents, "GET", nil},
		route{"/", addStudent, "POST", nil},
		route{"/{id}", getStudent, "GET", nil},
		route{"/{id}", updateStudent, "PUT", nil},
		route{"/{id}", deleteStudent, "DELETE", nil},
		route{"/attendance/{id}", punchIn, "POST", nil},
		route{"/attendance/{id}", punchOut, "PUT", nil},
		route{"/attendance/{id}", getAttendance, "GET",
			[]query{query{"start_date", "{start_date}"}, query{"end_date", "{end_date}"}},
		},
	}
}

func InitStudentsSubrouter(r *mux.Router) {
	s := r.PathPrefix("/students").Subrouter()

	routes := getRoutes()
	// register the routes on the above subrouter
	for _, rt := range routes {
		p := s.Path(rt.path).HandlerFunc(rt.handler).Methods(rt.method)
		for _, q := range rt.queries {
			p.Queries(q.key, q.value)
		}
	}
}
