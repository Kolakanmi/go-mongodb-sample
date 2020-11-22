package user

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) (string, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	UpdatePassword(ctx context.Context, id, password string) error
	Update(ctx context.Context, id string, user *model.User) error
	Delete(ctx context.Context, id string) error
}
