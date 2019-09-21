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

	// Route handlers
	// Static files (css, html, js, etc.)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("wwwroot/build"))))

	// Root handler /
	router.HandleFunc("/", controllers.Root)

	// If the cert files are present serve up an https server, redirect http to https
	if config.CertFile != "" && config.CertKeyFile != "" {

		httpServer := &http.Server{
			Addr: config.HTTP,
			Handler: http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) { // redirect to https
				target := "https://" + req.Host + req.URL.Path
				if len(req.URL.RawQuery) > 0 {
					target += "?" + req.URL.RawQuery
				}
				log.Printf("redirected to: %s", target)
				http.Redirect(res, req, target, http.StatusPermanentRedirect)
			}),
		}

		httpsServer := &http.Server{Addr: config.HTTPS, Handler: router}

		go httpServer.ListenAndServe()
		log.Fatal(httpsServer.ListenAndServeTLS(config.CertFile, config.CertKeyFile))

		// Otherwise serve up an http server
	} else {
		httpServer := &http.Server{Addr: config.HTTP, Handler: router}

		log.Fatal(httpServer.ListenAndServe())
	}
}
