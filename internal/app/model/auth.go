package model

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/uuid"
	"time"
)

type Auth struct {
	ID string `json:"id" bson:"_id"`
	CreatedAt *time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deleted_at"`
	UserID string `json:"userId" bson:"user_id"`
}

func (a *Auth) SetBase()  {
	created :=  timeNow()
	a.ID = uuid.New()
	a.CreatedAt = created
	a.UpdatedAt = created
	a.DeletedAt = nil
}
