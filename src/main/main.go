package main

import (
	"microservice-go-template/src/main/config"
	"microservice-go-template/src/main/service"
)

func main() {
	var cfg = config.GetConfig()
	service.Run(cfg)
}
