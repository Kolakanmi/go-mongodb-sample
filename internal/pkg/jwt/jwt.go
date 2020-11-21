package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type (
	Config struct {
		JWTSecretKey string `envconfig:"JWT_SECRET"`
		JWTIssuer string `envconfig:"JWT_ISSUER"`
		JWTExpiry time.Duration `envconfig:"JWT_EXPIRY"`
	}
	Claims struct {
		StandardClaims
		UserID string
		FirstName string
		LastName string
		Roles []string
	}
	Generator struct {
		Config Config
		SigningMethod jwt.SigningMethod
	}

	StandardClaims = jwt.StandardClaims
)

var ErrInvalidToken = errors.New("invalid token")

func New(conf Config) *Generator {
	return &Generator{
		Config:        conf,
		SigningMethod: jwt.SigningMethodHS256,
	}
}

func (g *Generator) Sign(claims Claims) (string, error) {
	token := jwt.NewWithClaims(g.SigningMethod, claims)
	v, err := token.SignedString([]byte(g.Config.JWTSecretKey))
	return v, err
}

func (g *Generator) Verify(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(g.Config.JWTSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}

