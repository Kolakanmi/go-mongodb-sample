package mongodb

import (
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
	_, _ = ConnectDB(conf)

	//fmt.Println(us)
}