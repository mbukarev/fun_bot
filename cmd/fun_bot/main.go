package main
import (
  "fmt"
  //"https://github.com/mbukarev/fun_bot/configs"
  "https://github.com/mbukarev/fun_bot/pkg/Config"
)

func main(){
  var cfg Config
  cfg.getConfig("fun_bot/configs/config.yaml")
  fmt.Print(cfg.tokenTelega)

}
