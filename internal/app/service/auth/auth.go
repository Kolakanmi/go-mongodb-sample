package auth

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
)

type Service struct {
	jwtSigner jwt.Signer
	authenticators map[string]IAuth
	authRepository auth.IAuthRepository
}

func NewService(signer jwt.Signer, repository auth.IAuthRepository) *Service {
	return &Service{
		jwtSigner: signer,
		authenticators: make(map[string]IAuth),
		authRepository: repository,
	}
}

func (s *Service) RegisterService(name string, authenticator IAuth)  {
	s.authenticators[name] = authenticator
}

func (s *Service) Auth(ctx context.Context, email, password string) (string, *model.User, error) {
	for _, a := range s.authenticators {
		user, err := a.Auth(ctx, email, password)
		if err != nil {
			log.WithContext(ctx).Error("failed to login with:", log.Field{
				"email": email,
				"error": err,
			})
			continue
		}
		newAuth := &model.Auth{
			UserID: user.ID,
		}
		authID, err := s.authRepository.Create(ctx, newAuth)
		if err != nil {
			log.WithContext(ctx).Error("failed to generate auth id ", log.Field{"err": err})
			return "", nil, err
		}
		token, err := s.jwtSigner.Sign(userAndAuthIDToClaims(user, authID))
		if err != nil {
			log.WithContext(ctx).Error("failed to generate jwt: ", log.Field{
				"error": err,
			})
			return "", nil, err
		}
		return token, user, nil
	}
	return "", nil, apperror.BadRequestError("Invalid user details")
}
