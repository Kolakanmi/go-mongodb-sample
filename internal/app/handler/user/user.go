package user

import (
	"encoding/json"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user/reqresponse"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"net/http"
)

type Handler struct {
	service user.IUserService
}

func NewHandler(srv user.IUserService) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) error {
	var req reqresponse.RequestRegister
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return apperror.BadRequestError("Invalid input")
	}
	defer r.Body.Close()
	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		return apperror.UserFriendlyError("Could not complete request", 500)
	}
	return response.OK("Success", user.Strip()).ToJSON(w)
}

func (h *Handler) FindAll(w http.ResponseWriter, r *http.Request) error {
	users, err := h.service.FindAll(r.Context())
	if err != nil {
		return apperror.UserFriendlyError("Could not complete request", 500)
	}
	return response.OK("Success", users).ToJSON(w)
}
