package user

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/handler"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/router"
	"net/http"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/users",
			Method:      http.MethodGet,
			Handler:     handler.Handler(h.FindAll),
			Middlewares: []router.Middleware{auth.RequiredAuthMiddleware([]string{"user"})},
				},
		{
			Path:        "/api/v1/register",
			Method:      http.MethodPost,
			Handler:     handler.Handler(h.Register),
				},
	}
}
