package auth

import (
	"encoding/json"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/auth/reqresponse"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"net/http"
)

type Handler struct {
	service auth.IService
}

func NewHandler(srv auth.IService) *Handler {
	return &Handler{service: srv}
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) error {
	req := reqresponse.ReqLogin{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	token, user, err := h.service.Auth(r.Context(), req.Email, req.Password)
	if err != nil {
		return err
	}
	result := reqresponse.ResponseLogin{
		Token: token,
		User:  *user,
	}
	return response.OK("success", result).ToJSON(w)
}