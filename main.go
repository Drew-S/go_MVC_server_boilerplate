package main

import (
	"log"
	"net/http"

	"./controllers"
	"./utils"

	"github.com/gorilla/mux"
)

var config utils.Config = utils.GetConfig()

func main() {

	router := mux.NewRouter()

	// Router handlers
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("wwwroot/build"))))

	router.HandleFunc("/", controllers.Root)

	// Run server
	server := &http.Server{
		Addr:    config.Addr,
		Handler: router,
	}

	if config.CertFile != "" && config.CertKeyFile != "" {
		log.Fatal(server.ListenAndServeTLS(config.CertFile, config.CertKeyFile))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
