package auth

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/handler"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/router"
	"net/http"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path: "/api/v1/auth",
			Method: http.MethodPost,
			Handler: handler.Handler(h.Auth),
		},
	}
}
