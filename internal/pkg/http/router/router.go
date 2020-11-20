package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type (
	//Instantiate it to avoid casting
	Middleware = func(handler http.Handler) http.Handler

	Route struct {
		Path string
		Method string
		Queries []string
		Handler http.Handler
		Middlewares []Middleware
	}

	Config struct {
		GlobalMiddlewares []Middleware
		Routes []Route

	}
)

func New(conf *Config) (http.Handler, error)  {
	r := mux.NewRouter()
	for _, middleware := range conf.GlobalMiddlewares {
		r.Use(middleware)
	}
	for _, route := range conf.Routes {
		h := route.Handler
		//Last middleware will be the innermost middleware and executed first.
		for i := len(route.Middlewares) -1; i >= 0; i-- {
			h = route.Middlewares[i](h)
		}
		r.Path(route.Path).Methods(route.Method).Handler(h).Queries(route.Queries...)
	}
	return r, nil
}

func GetEmptyConfig() *Config {
	return &Config{}
}