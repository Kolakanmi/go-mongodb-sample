package handler

import (
	"errors"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/apperror"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/response"
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := h(writer, request)
	if err == nil {
		return
	}
	err = respondWithError(err, writer)
	if err != nil {
		log.Println(err)
	}
}


func respondWithError(err error, w http.ResponseWriter) error {
	appError := new(apperror.AppError)
	if errors.As(err, &appError) {
		return response.Fail(appError, appError.Type()).ToJSON(w)
	}
	return response.Fail(apperror.ErrInternalServer, http.StatusInternalServerError).ToJSON(w)
}