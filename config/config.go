package config

import (
	"encoding/json"
	"os"
)

//Config ...
type Config struct {
	MariaDB `json:"marialdb"`
}

//LoadConfiguration ...
func LoadConfiguration(file string) Config {
	var err error
	var config Config
	var jsonParser *json.Decoder
	var configFile *os.File

	configFile, err = os.Open(file)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	jsonParser = json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		panic(err)
	}
	return config
}
