package middleware

import "net/http"

func SetResponseHeaders(h http.Handler) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "")
		//insert other headers
		h.ServeHTTP(w, r)
	})
}
