package golaxy

import (
	"os"
	"encoding/json"
	"github.com/mskoroglu/golaxy/http"
)

type Properties struct {
	Server struct {
		IP   string `json:"ip"`
		Port int `json:"port"`
	} `json:"server"`
}

var props = Properties{}

var Run = func(args ...interface{}) {
	loadProperties()
	http.StartHttpServer(props.Server.IP, props.Server.Port)
}

func loadProperties() {
	configFile, err := os.Open("app.json")
	if err != nil {
		println(err.Error())
	}

	parser := json.NewDecoder(configFile)
	if err = parser.Decode(&props); err != nil {
		println(err.Error())
	}
}
