package mongodb

import (
	"context"
	"fmt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Config struct {
	Username string `envconfig:"MONGODB_USERNAME"`
	Password string `envconfig:"MONGODB_PASSWORD"`
	Database string `envconfig:"MONGODB_DATABASE"`
	Address string `envconfig:"MONGODB_ADDRESS"`
}

func LoadConfigFromFile() *Config {
	var config Config
	envconfig.Load(&config)
	return &config
}

func ConnectDB(conf *Config) (*mongo.Database, error){
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.fsldx.mongodb.net/%s?retryWrites=true&w=majority", conf.Username,
		conf.Password, conf.Database)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	database := client.Database(conf.Database)
	return database, err
}