package main

import (
	"github.com/Gircen/go-library-core/main/config"
	"microservice-go-template/src/main/service"
)

func main() {
	var cfg = config.GetConfig()
	service.Run(cfg)
}
