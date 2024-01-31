package middleware

import "net/http"

func Corsmw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    // we need to allow this here
    // even ports make a difference for the origin
    // https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
    // If we are using the Authorization header we need to specify this here
    // https://stackoverflow.com/questions/10548883/request-header-field-authorization-is-not-allowed-error-tastypie
    w.Header().Set("Access-Control-Allow-Headers", "Authorization")
    if r.Method == http.MethodOptions {
      return
    }

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
