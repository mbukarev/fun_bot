package config
import (
  "os"
  "github.com/go-yaml/yaml"

)
type Config struct{
  token string `envconfig:"tokenTelega"`
}

func (c *Config) getConfig (file string) {
  f, err := os.Open(file)
  if err != nil {
      processError(err)
  }
  defer f.Close()

  var cfg Config
  decoder := yaml.NewDecoder(f)
  err = decoder.Decode(&c)
  if err != nil {
      processError(err)
  }
}
