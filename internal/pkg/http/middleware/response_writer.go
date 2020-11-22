package middleware

import "net/http"

type MyResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *MyResponseWriter) WriteHeader(code int)  {
	w.ResponseWriter.WriteHeader(code)
	w.status = code
}

func (w *MyResponseWriter) Status() int {
	return w.status
}
//StatusResponseWriter - ResponseWriters are passed as a copy to each handler. This handler wraps the ResponseWriter
//into a pointer type that implements the ResponseWriter interface.
func StatusResponseWriter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw := &MyResponseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		h.ServeHTTP(mw, r)
	})
}
