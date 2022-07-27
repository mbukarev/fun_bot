package Config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"os"
)

type Config struct {
	Token string `yaml:"TokenTelega"`
}

func (c *Config) GetConfig(file string) {
	f, err := os.Open(file)
	if err != nil {
		processError(err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	fmt.Println(err)
}
