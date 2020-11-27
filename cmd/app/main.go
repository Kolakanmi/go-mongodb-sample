package main

import (
	"github.com/Kolakanmi/go-mongodb-sample/internal/app/api"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/http/server"
	log2 "github.com/Kolakanmi/go-mongodb-sample/internal/pkg/log"
	"log"
)

func main() {
	err := envconfig.SetEnvFromFile("config.env")
	if err != nil {
		log.Println("envconfig load err: ", err)
	}
	lConf := log2.LoadEnvConfig()
	log2.SetLogger(*lConf)
	log2.Init(log2.Field{"service": "mongo-sample"})
	log2.Info("Initializing http routing...", log2.Field{})
	router, err := api.NewRouter()
	if err != nil {
		log.Println(err)
	}
	serverConf := server.LoadConfigFromEnv()
	server.ListenAndServe(*serverConf, router)
}
