package auth

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"
	"time"
)

func NewContext(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, authContextKey, user)
}

func FromContext(ctx context.Context) *model.User {
	v, ok := ctx.Value(authContextKey).(*model.User)
	if ok {
		return v
	}
	return nil
}

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

func claimsToUserAndAuthID(claims *jwt.Claims) (*model.User, string) {
	return &model.User{
		ID: claims.UserID,
		FirstName: claims.FirstName,
		LastName:  claims.LastName,
		Email:     claims.Email,
		Roles: claims.Roles,
	},
	claims.Id
}
