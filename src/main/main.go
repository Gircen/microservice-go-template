package main

import (
	"github.com/Gircen/go-library-core/main/config"
	"github.com/Gircen/go-library-core/main/logs"
	"github.com/Gircen/go-library-core/main/service"
	"microservice-go-template/src/main/service/crud"
)

func main() {

	var cfg = config.GetConfig()
	service.RunServiceCore(cfg)
	logs.Logger.Println("ERRRRR")
	crud.Run(cfg)
}
