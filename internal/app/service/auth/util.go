package auth

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"
	"time"
)

func userAndAuthIDToClaims(user *model.User, authID string) jwt.Claims {
	return jwt.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwt.JWTExpiry).Unix(),
			Id:        authID,
			IssuedAt:  time.Now().Unix(),
			Issuer:    jwt.JWTIssuer,
			Subject:   user.ID,
		},
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email: user.Email,
		Roles:     user.Roles,
	}
}

func claimsToUserAndAuthID(claims jwt.Claims) (*model.User, string) {
	return &model.User{
		Base:      model.Base{
			ID: claims.UserID,
		},
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Email:     claims.Email,
		Roles: claims.Roles,
	},
	claims.Id
}
