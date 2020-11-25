package user

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user/reqresponse"
)

type IUserService interface {
	auth.IAuth
	Register(ctx context.Context, req reqresponse.RequestRegister) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}
