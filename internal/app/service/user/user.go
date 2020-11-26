package user

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/user"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/service/user/reqresponse"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/db/mongodb"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo user.IUserRepository
}

func NewUserService(repo user.IUserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Auth(ctx context.Context, email, password string) (*model.User, error) {
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

func (s *Service) Register(ctx context.Context, req reqresponse.RequestRegister) (*model.User, error) {
	existing, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil && !mongodb.IsErrNotFound(err) {
		log.WithContext(ctx).Error("failed to check use by email", log.Field{"err": err})
		return nil, apperror.UserFriendlyError("could not complete your request", 500)
	}
	if existing != nil {
		log.WithContext(ctx).Error("user already exists", log.Field{"err": err})
		return nil, apperror.BadRequestError("user already exists")
	}
	password, err := s.generatePassword(req.Password)
	if err != nil {
		log.WithContext(ctx).Error("error generating bcrypt password", log.Field{"err": err})
		return nil, apperror.UserFriendlyError("could not complete your request", 500)
	}
	user := &model.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  password,
		Roles:     []string{"user"},
	}
	_, err = s.repo.Create(ctx, user)
	if err != nil {
		log.WithContext(ctx).Error("failed to insert user", log.Field{"err": err})
		return nil, apperror.UserFriendlyError("could not complete request", 500)
	}
	return user.Strip(), nil

}

func (s *Service) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		log.WithContext(ctx).Error("failed to insert user", log.Field{"err": err})
		return "", apperror.UserFriendlyError("could not complete request", 500)
	}
	return id, nil
}

func (s *Service) FindByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if mongodb.IsErrNotFound(err) {
			log.WithContext(ctx).Error("user with id does not exist", log.Field{"id": id,"err": err})
			return nil, apperror.UserFriendlyError("could not complete request", 500)
		}
		log.WithContext(ctx).Error("error fetching user", log.Field{"id": id,"err": err})
		return nil, apperror.UserFriendlyError("could not complete request", 500)
	}
	return user.Strip(), nil
}

func (s Service) FindAll(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		if mongodb.IsErrNotFound(err) {
			log.WithContext(ctx).Error("user with id does not exist", log.Field{"err": err})
			return nil, apperror.UserFriendlyError("could not complete request", 500)
		}
		log.WithContext(ctx).Error("error fetching user", log.Field{"err": err})
		return nil, apperror.UserFriendlyError("could not complete request", 500)
	}
	strippedUsers := make([]*model.User, 0)
	for _,v := range users {
		strippedUsers = append(strippedUsers, v.Strip())
	}
	return strippedUsers, nil
}

func (s Service) Update(ctx context.Context, user *model.User) error {
	err := s.repo.Update(ctx, user.ID, user)
	if err != nil {
		log.WithContext(ctx).Error("error updating user", log.Field{"id": user.ID,"err": err})
		return apperror.UserFriendlyError("could not complete request", 500)
	}
	return nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		log.WithContext(ctx).Error("error deleting user", log.Field{"id": id,"err": err})
		return apperror.UserFriendlyError("could not complete request", 500)
	}
	return nil
}

func (s *Service) generatePassword(password string) (string, error)  {
	p, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(p), err
}