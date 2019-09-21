package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Config struct used with config.json
type Config struct {
	HTTP        string `json:"http"`
	HTTPS       string `json:"https"`
	CertFile    string `json:"ssl_cert_file"`
	CertKeyFile string `json:"ssk_cert_key_file"`
	SessionKey  string `json:"session_key"`
	SQL         string `json:"sql_connect"`
	SQLDriver   string `json:"sql_driver"`
	Dev         bool   `json:"dev"`
	Name        string `json:"name"`
}

var config Config
var loaded bool = false

// GetConfig returns a Config object created from the config.json file
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
