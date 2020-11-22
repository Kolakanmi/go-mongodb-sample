package reqresponse

type (

	RequestRegister struct {
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	RequestGeneratePasswordResetToken struct {
		Email string `json:"email"`
	}

	RequestUpdatePassword struct {
		Token string `json:"token"`
		Password string `json:"password"`
	}
)
