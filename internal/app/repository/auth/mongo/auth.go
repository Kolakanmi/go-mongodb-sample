package mongo

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	database *mongo.Database
}

func NewAuthRepository(database *mongo.Database) *AuthRepository  {
	return &AuthRepository{database: database}
}

func (a *AuthRepository) Create(ctx context.Context, auth *model.Auth) (string, error) {
	_ = a.DeleteByUserID(ctx, auth.UserID)
	auth.Base = model.SetBase()
	_, err := a.collection().InsertOne(ctx, auth)
	if err != nil {
		return "", err
	}
	return auth.ID, nil
}

func (a *AuthRepository) FindByUserID(ctx context.Context, userID string) (*model.Auth, error) {
	filter := bson.M{"user_id": userID}
	var auth *model.Auth
	err := a.collection().FindOne(ctx, filter).Decode(auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *AuthRepository) FindByID(ctx context.Context, id string) (*model.Auth, error) {
	filter := bson.M{"_id": id}
	var auth *model.Auth
	err := a.collection().FindOne(ctx, filter).Decode(auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (a *AuthRepository) DeleteByID(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := a.collection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) DeleteByUserID(ctx context.Context, userID string) error {
	filter := bson.M{"user_id": userID}
	_, err := a.collection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) collection() *mongo.Collection {
	return a.database.Collection("auth")
}
