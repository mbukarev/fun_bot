package main

import (
	. "fun_bot/pkg/Config"
)

func main() {
	var cfg Config
	cfg.GetConfig("configs/config.yaml")
	
}
