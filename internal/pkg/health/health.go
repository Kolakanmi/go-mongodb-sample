package health

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/handler"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"net/http"
	"sync"
)

var (
	isReadyMutex sync.RWMutex
	isReady bool
)

func Ready()  {
	isReadyMutex.Lock()
	isReady = true
	isReadyMutex.Unlock()
}
//Returns an handler for checking the readiness stste
func Readiness() http.Handler {
	return handler.Handler(func(w http.ResponseWriter, r *http.Request) error {
		isReadyMutex.RLock()
		defer isReadyMutex.RUnlock()
		if !isReady {
			return response.Fail(apperror.UserFriendlyError("Server not ready", 503)).ToJSON(w)
		}
		return response.OK("Server is up!", nil).ToJSON(w)
	})
}