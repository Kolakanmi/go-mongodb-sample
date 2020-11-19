package mongodb

import (
	"fmt"
	"github.com/Kolakanmi/go-mongodb-sample/internal/pkg/envconfig"
	"testing"
)

func TestMongoConnection(t *testing.T) {
	err := envconfig.SetEnvFromFile("")
	if err != nil {
		t.Error(err)
	}
	conf := LoadConfigFromFile()
	cl, cancel, err := ConnectDB(conf)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("Connected! ", cl)
	cancel()
}