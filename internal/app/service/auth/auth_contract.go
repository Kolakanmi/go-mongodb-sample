package auth

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
)

type IAuth interface {
	Auth(ctx context.Context, email, password string) (*model.User, error)
}

type IService interface {
	Auth(ctx context.Context, email, password string) (string, *model.User, error)
}

