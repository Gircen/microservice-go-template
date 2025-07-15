package service

import (
	"log"
	"microservice-go-template/src/main/config"
	"net/http"
)

func Run(config *config.Config) {
	http.HandleFunc("/healthCheck", HealthCheck)
	log.Printf(config.Server.Host + ":" + config.Server.Port)
	err := http.ListenAndServe(config.Server.Host+":"+config.Server.Port, nil)
	if err != nil {
		println(err)
	}
}
