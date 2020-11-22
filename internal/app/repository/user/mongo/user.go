package mongo

import (
	"context"
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database *mongo.Database
}

func NewUserRepository(database *mongo.Database)  *UserRepository{
	return &UserRepository{database: database}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) (string, error) {
	user.Base = model.SetBase()
	_, err := r.collection().InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	filter := bson.M{"_id": id}
	var user *model.User
	err := r.collection().FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	cursor, err := r.collection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user *model.User
		err := cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	filter := bson.M{"email": email}
	var user *model.User
	err := r.collection().FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id, password string) error {
	panic("implement me")
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (r *UserRepository) collection() *mongo.Collection {
	return r.database.Collection("user")
}
