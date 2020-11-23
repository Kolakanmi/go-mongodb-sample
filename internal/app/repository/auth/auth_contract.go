package auth

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
)

type IAuthRepository interface {
	Create(ctx context.Context, auth *model.Auth) (string, error)
	FindByUserID(ctx context.Context, userID string) (*model.Auth, error)
	FindByID(ctx context.Context, id string) (*model.Auth, error)
	DeleteByID(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
}
