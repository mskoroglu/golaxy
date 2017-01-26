package config

import (
	"os"
	"encoding/json"
)

type Properties struct {
	Server struct {
		IP   string `json:"ip"`
		Port int `json:"port"`
	} `json:"server"`

	DataSource struct {
		Sql struct {
			Driver string `json:"driver"`
			Url    string `json:"url"`
		} `json:"sql"`
	} `json:"data_source"`
}

var props *Properties

func GetProperties() *Properties {
	if props == nil {
		file := getConfigurationFile()
		if file == nil {
			return nil
		}
		parser := json.NewDecoder(file)
		decodeError := parser.Decode(&props)
		if decodeError != nil {
			println(decodeError.Error())
		}
	}
	return props
}

func getConfigurationFile() *os.File {
	configFile, err := os.Open("app.json")
	if err != nil {
		return nil
	}
	return configFile
}
