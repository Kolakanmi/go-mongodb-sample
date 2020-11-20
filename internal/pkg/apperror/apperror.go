package apperror

import "errors"

type (

	AppError struct {
		text string
		status int
	}
)

var ErrInternalServer = errors.New("nternal server error")

func (a AppError) Error() string{
	return a.text
}

func (a AppError) Type() int {
	return a.status
}

func BadRequestError(text string) error {
	return AppError{
		text:    text,
		status: 400,
	}
}

func NotFound(entity string) error {
	return AppError{
		text:    entity + "not found",
		status: 404,
	}
}

func UserFriendlyError(text string, status int)  error {
	return AppError{
		text:    text,
		status: status,
	}
}

func IsAppError(err error) bool {
	appError := new(AppError)
	if errors.As(err, &appError) {
		return true
	}
	return false
}
