package handler

import (
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
	return response.Fail(err).ToJSON(w)
}