package reqresponse

import "github.com/Kolakanmi/go-mongodb-sample/internal/app/model"

type ResponseLogin struct {
	Token string `json:"token"`
	User model.User `json:"user"`
}
