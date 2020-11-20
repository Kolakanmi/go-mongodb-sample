package server

import (
	"context"
	"fmt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type (
	Config struct {
		Address string `envconfig:"HTTP_ADDRESS"`
		Port string `envconfig:"HTTP_PORT"`
		ReadTimeout time.Duration `envconfig:"HTTP_READ_TIMEOUT"`
		ReadHeaderTimeout time.Duration `envconfig:"HTTP_READ_HEADER_TIMEOUT"`
		WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT"`
		IdleTimeout time.Duration `envconfig:"HTTP_IDLE_TIMEOUT"`
		ShutdownTimeout time.Duration `envconfig:"HTTP_SHUTDOWN_TIMEOUT"`
	}
)

func ListenAndServe(conf Config, handler http.Handler)  {
	port := conf.Port
	if  port == "" {
		appEnginePort := os.Getenv("PORT")
		if appEnginePort == "" {
			port = "80"
		}
	}
	address := fmt.Sprintf("%s:%s", conf.Address, port)

	srv := http.Server{
		Addr: address,
		Handler: handler,
		ReadTimeout: conf.ReadTimeout,
		ReadHeaderTimeout: conf.ReadHeaderTimeout,
		WriteTimeout: conf.WriteTimeout,
		IdleTimeout: conf.IdleTimeout,
	}

	log.Println("Listening on: " + address)
	//Run server in goroutine to implement graceful shutdown.
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("Server listen and serve error: ", err)
		}
	}()
	//Graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	//Block until signal is received
	<- signals

	ctx, cancel := context.WithTimeout(context.Background(), conf.ShutdownTimeout * time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Println("Shutdown error")
	}

}

func LoadConfigFromEnv() *Config {
	var config Config
	envconfig.Load(&config)
	return &config
}
