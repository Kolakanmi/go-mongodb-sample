package handler

import (
	"log"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) error

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := h(writer, request)
	if err == nil {
		return
	}
	log.Println(err)
}
