package user

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/user"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user/reqresponse"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo user.IUserRepository
}

func (s Service) Auth(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		log.WithContext(ctx).Error("failed to check existing user by email", log.Field{"err": err})
		return nil, apperror.UserFriendlyError("invalid email or password", 401)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.WithContext(ctx).Error("invalid password", log.Field{})
		return nil, apperror.UserFriendlyError("invalid email or password", 401)
	}
	return user.Strip(), nil
}

func (s Service) Register(ctx context.Context, req reqresponse.RequestRegister) (*model.User, error) {
	panic("implement me")
}

func (s Service) Create(ctx context.Context, user *model.User) (string, error) {
	panic("implement me")
}

func (s Service) FindByID(ctx context.Context, id string) (*model.User, error) {
	panic("implement me")
}

func (s Service) FindAll(ctx context.Context) ([]*model.User, error) {
	panic("implement me")
}

func (s Service) ResetPassword(ctx context.Context, req reqresponse.RequestGeneratePasswordResetToken) error {
	panic("implement me")
}

func (s Service) Update(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (s Service) UpdatePassword(ctx context.Context, req reqresponse.RequestUpdatePassword) {
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
