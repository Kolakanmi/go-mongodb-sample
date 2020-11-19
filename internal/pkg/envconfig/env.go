package envconfig

import (
	"bufio"
	"errors"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
	"strings"
)

var envPrefix = ""

func Load(t interface{})  {
	LoadWithPrefix(envPrefix, t)
}

func LoadWithPrefix(prefix string, t interface{}) {
	if err := envconfig.Process(prefix, t); err != nil {
		log.Println("Unable to read")
	}
}
func SetEnvFromFile(f string) error {
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#") || strings.TrimSpace(txt) == "" {
			continue
		}
		env := strings.SplitN(txt, "=", 2)
		if len(env) != 2 {
			return errors.New("must be a key-value pair")
		}
		key := env[0]
		value := env[1]
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return nil
}