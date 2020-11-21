package mongodb

import (
	"context"
	"fmt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"testing"
)

type MyTest struct {
	ID string `bson:"_id"`
	Name string `bson:"name"`
}

func TestMongoConnection(t *testing.T) {
	err := envconfig.SetEnvFromFile("")
	if err != nil {
		t.Error(err)
	}
	conf := LoadConfigFromFile()
	client, cancel, err := ConnectDB(conf)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("Connected! ", client)
	collection := client.Database("testing").Collection("test")
	my := MyTest{
		ID: "111111",
		Name: "TEST",
	}
	i, err := collection.InsertOne(context.Background(), my)
	if err != nil {
		t.Error(err)
	}
	
	fmt.Println("i", i)
	cancel()
}