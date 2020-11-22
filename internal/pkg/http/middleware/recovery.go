package middleware

import (
	"fmt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"net/http"
	"runtime"
)

func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request){
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", err)
				}
				stack := make([]byte, 4<<10) //4kB
				length := runtime.Stack(stack, false)
				log.WithContext(r.Context()).Error("Panic recovery", log.Field{"err": err, "stack": string(stack[:length])})
				_ = response.Fail(apperror.ErrInternalServer, 500).ToJSON(w)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
