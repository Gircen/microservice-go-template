package main

import (
	"github.com/Gircen/go-library-core/main/logs"
	"github.com/Gircen/go-library-core/main/service"
	"microservice-go-template/src/main/service/crud"
)

func main() {

	service.Init()
	logs.Info("start")
	crud.Run()
}
