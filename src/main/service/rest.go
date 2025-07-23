package service

import (
	"github.com/Gircen/go-library-api/main/config"
	"github.com/Gircen/go-library-core/main/service"
	"log"
	"net/http"
)

func Run(config *config.Config) {
	http.HandleFunc("/healthCheck", service.HealthCheck)
	log.Printf(config.Server.Host + ":" + config.Server.Port)
	err := http.ListenAndServe(config.Server.Host+":"+config.Server.Port, nil)
	if err != nil {
		println(err)
	}
}
