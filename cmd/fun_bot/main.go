package main
import (
  "fmt"
  //"https://github.com/mbukarev/fun_bot/configs"
  "fun_bot/pkg/Config"
)

func main(){
  var cfg Config
  cfg.getConfig("../configs/config.yaml")
  fmt.Print(cfg.tokenTelega)

}
