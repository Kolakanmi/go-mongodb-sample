package auth

import "github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"

type Service struct {
	jwtSigner jwt.Signer
	authenticators map[string]IAuth
}

func NewService(signer jwt.Signer) *Service {
	return &Service{
		jwtSigner: signer,
		authenticators: make(map[string]IAuth),
	}
}

func (s *Service) RegisterService(name string, authenticator IAuth)  {
	s.authenticators[name] = authenticator
}