package config
import (
  "os"
  "yaml"

)
type Config strcut{
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
