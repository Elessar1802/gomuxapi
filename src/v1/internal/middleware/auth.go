package middleware

import "net/http"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if not authenticated user then we cannot use

		next.ServeHTTP(w, r)
	})
}
