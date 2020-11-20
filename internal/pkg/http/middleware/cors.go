package middleware

import (
	"github.com/gorilla/handlers"
	"net/http"
)

func Cors(h http.Handler) http.Handler {
	headersOK := handlers.AllowedHeaders([]string{"Content-Type", "Accept", "Authorization"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	c := handlers.CORS(headersOK, originsOK, methodsOK)
	return c(h)
	//Or Simply return handlers.CORS(headersOK, originsOK, methodsOK)(h)
}
