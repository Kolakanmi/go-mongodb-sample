package model

import "time"

type Base struct {
	ID string `json:"id" bson:"_id"`
	CreatedAt *time.Time `json:"createdAt" bson:"created_at"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deleted_at"`
}
