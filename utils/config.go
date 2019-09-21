package utils

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"path"
)

type Config struct {
	Addr        string `json:"addr"`
	Port        string `json:"port"`
	CertFile    string `json:"ssl_cert_file"`
	CertKeyFile string `json:"ssk_cert_key_file"`
	SessionKey  string `json:"session_key"`
	SQL         string `json:"sql_connect"`
	SQLDriver   string `json:"sql_driver"`
	Dev         bool   `json:"dev"`
	Name        string `json:"name"`
}

func CreateTemplate(name string) (*template.Template, error) {
	return template.New(path.Base(name)).
		Funcs(template.FuncMap{
			"title": func() string {
				if GetConfig().Name == "" {
					return "Site name"
				}
				return GetConfig().Name
			},
			"dev": func() bool {
				return GetConfig().Dev
			},
		}).
		ParseFiles([]string{
			name,
			"views/shared/layout.html",
		}...)
}

type Template struct {
	Title string
	Dev   bool
	Data  interface{}
}

var config Config
var loaded bool = false

func GetConfig() Config {
	if !loaded {
		byteData, err := ioutil.ReadFile("config.json")
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(byteData, &config)
		loaded = true
	}
	return config
}
