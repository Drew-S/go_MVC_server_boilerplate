package controllers

import (
	"net/http"

	"../models"
	"../utils"
)

// Root handles connections to the path /
func Root(res http.ResponseWriter, req *http.Request) {
	tmpl, err := utils.CreateTemplate("views/root/index.html")
	if err != nil {
		res.Write([]byte("An error occurred in rendering the template"))
		return
	}

	// db := models.GetDatabase()

	var front []models.Example = []models.Example{
		models.Example{Name: "bootstrap", URL: "https://getbootstrap.com"},
	}

	var server []models.Example = []models.Example{
		models.Example{Name: "gorilla/mux", URL: "https://github.com/gorilla/mux"},
		models.Example{Name: "gorilla/sessions", URL: "https://github.com/gorilla/sessions"},
		models.Example{Name: "go-gorp/gorp", URL: "https://github.com/go-gorp/gorp"},
	}

	// db.Select(&data, "select * from example Id")

	err = tmpl.Execute(res, struct {
		Title  string
		Server []models.Example
		Front  []models.Example
	}{
		Title:  "root",
		Server: server,
		Front:  front,
	})
	if err != nil {
		res.Write([]byte(err.Error()))
	}
}
