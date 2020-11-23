package model

type Auth struct {
	Base
	UserID string `json:"userId" bson:"user_id"`
}
