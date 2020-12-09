package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

//Config ...
type Config struct {
	MariaDB `json:"mariadb"`
}

//LoadConfiguration ...
func LoadConfiguration(file string) Config {
	var err error
	var config Config
	var jsonParser *json.Decoder
	var configFile *os.File
	var workDir string

	workDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	configFile, err = os.Open(filepath.Join(workDir, file))
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
