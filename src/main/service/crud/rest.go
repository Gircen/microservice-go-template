package crud

import (
	"fmt"
	"github.com/Gircen/go-library-core/main/conf"
	"github.com/Gircen/go-library-core/main/logs"
	"net/http"
)

var HttpServer = http.NewServeMux()

func GetService(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/get/hi" {

		if req.Method == http.MethodGet {
			_, err := fmt.Fprint(w, "Hi")
			if err != nil {
				return
			}
		}
	}
}

func PostService(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path == "/post/hi" {

		if req.Method == http.MethodPost {
			_, err := fmt.Fprint(w, "Hi")
			if err != nil {
				return
			}
		}
	}
}

func Run() {
	HttpServer.HandleFunc("/get/hi", GetService)
	HttpServer.HandleFunc("/post/hi", PostService)
	logs.Debug(conf.Conf.Server.Host + ":" + conf.Conf.Server.Port)
	err := http.ListenAndServe(conf.Conf.Server.Host+":"+conf.Conf.Server.Port, HttpServer)
	if err != nil {
		logs.FatalError(err)
	}
}
