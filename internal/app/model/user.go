package model

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/uuid"
	"time"
)

type (
	User struct {
		ID string `json:"id" bson:"_id"`
		CreatedAt *time.Time `json:"createdAt" bson:"created_at"`
		UpdatedAt *time.Time `json:"updatedAt" bson:"updated_at"`
		DeletedAt *time.Time `json:"deletedAt" bson:"deleted_at"`
		FirstName string `json:"firstName" bson:"first_name"`
		LastName string `json:"lastName" bson:"last_name"`
		Email string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
		Roles []string `json:"roles" bson:"roles"`
	}
)

func (u *User) SetBase()  {
	created :=  timeNow()
	u.ID = uuid.New()
	u.CreatedAt = created
	u.UpdatedAt = created
	u.DeletedAt = nil
}