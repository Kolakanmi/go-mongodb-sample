package auth

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/repository/auth"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/jwt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"net/http"
	"strings"
)

type contextKey string

const authContextKey = "auth_key"

func UserAuthMiddleware(verifier jwt.Verifier, authRepo auth.IAuthRepository) func(handler http.Handler) http.Handler  {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")
			if bearerToken == "" {
				//For routes that do not need authentication
				next.ServeHTTP(w, r)
				return
			}
			bearerPrefix := "Bearer "
			token := ""
			if strings.HasPrefix(bearerToken, bearerPrefix) {
				if len(bearerToken) > len(bearerPrefix) {
					token = bearerToken[len(bearerPrefix):]
				} else {
					log.WithContext(r.Context()).Error("empty token", log.Field{})
				}
			} else {
				log.WithContext(r.Context()).Info("Not in bearer token format", log.Field{})
				token = bearerToken
			}
			token = strings.TrimSpace(token)

			claims, err := verifier.Verify(token)
			if err != nil {
				log.WithContext(r.Context()).Error("Invalid jwt ", log.Field{"err": err})
				next.ServeHTTP(w, r)
				return
			}

			user, authID := claimsToUserAndAuthID(claims)

			_, err = authRepo.FindByID(r.Context(), authID)
			if err != nil {
				log.WithContext(r.Context()).Error("token blacklisted ", log.Field{"err": err})
				next.ServeHTTP(w, r)
				return
			}
			newContext := NewContext(r.Context(), user)
			r = r.WithContext(newContext)
			next.ServeHTTP(w,r)
			return
		})
	}
}

func RequiredAuthMiddleware(roles []string) func(http.Handler) http.Handler  {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := FromContext(r.Context())
			if user == nil {
				log.WithContext(r.Context()).Error("unauthorized", log.Field{})
				response.FailWriter(w, apperror.UserFriendlyError("Unauthorized", 401))
				return
			}
			if len(roles) > 0 && len(user.Roles) == 0 {
				log.WithContext(r.Context()).Error("unauthorized", log.Field{})
				response.FailWriter(w, apperror.UserFriendlyError("Unauthorized", 401))
				return
			}
			for _, authRole := range roles {
				for _, userRole := range user.Roles {
					if authRole == userRole {
						next.ServeHTTP(w,r)
						return
					}
				}
			}
			response.FailWriter(w, apperror.UserFriendlyError("Unauthorized", 401))
		})
	}
}
