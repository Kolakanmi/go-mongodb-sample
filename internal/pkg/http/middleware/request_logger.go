package middleware

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"net/http"
	"strings"
	"time"
)

func HTTPRequestResponseInfo(ignorePrefix []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, ignore := range ignorePrefix {
				if strings.HasPrefix(r.URL.Path, ignore) {
					next.ServeHTTP(w, r)
					return
				}
			}
			timeNow := time.Now()
			log.WithContext(r.Context()).With(log.Field{"stage": "request"}).Info("started", log.Field{})

			next.ServeHTTP(w, r)

			responseCode := 200

			if mw, ok := w.(interface{
				Status() int
			}); ok {
				responseCode = mw.Status()
			}

			responseTime := time.Since(timeNow)
			log.WithContext(r.Context()).With(log.Field{
				"status": responseCode,
				"response_time": responseTime,
				"stage": "response",
			}).Info("request finished", log.Field{})
		})
	}
}
