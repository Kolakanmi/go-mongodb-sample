package model

type (
	User struct {
		Base
		FirstName string `json:"firstName" bson:"first_name"`
		LastName string `json:"lastName" bson:"last_name"`
		Email string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
		Roles []string `json:"roles" bson:"roles"`
	}
)
