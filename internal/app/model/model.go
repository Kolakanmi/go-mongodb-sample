package model

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/uuid"
	"time"
)

type Base struct {
	ID string `json:"id" bson:"_id"`
	CreatedAt *time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deleted_at"`
}

func SetBase() Base {
	created :=  timeNow()
	return Base{
		ID:        uuid.New(),
		CreatedAt: created,
		UpdatedAt: created,
		DeletedAt: nil,
	}
}

func timeNow() *time.Time {
	t := time.Now()
	return &t
}
