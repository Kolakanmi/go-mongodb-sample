package log

import "net/http"

//NewHTTPContextHandler adds a context logger to each request
func NewHTTPContextHandler(l Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := l.With(Field{
				"path": r.URL.Path,
				"method": r.Method,
			})
			r = r.WithContext(NewContext(ctx, logger))
			next.ServeHTTP(w, r)
		})
	}
}
