package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"tspdev.com/m/v2/config"
)

func main() {
	obj, err := getDBConfig()
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}

	fmt.Println(obj)
}

func getDBConfig() (*config.DBConfig, error) {
	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, err
	}
	config_path := directory + "/config/dbconfig.yaml"
	_, err = os.Stat(config_path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New("Could not find config file at: " + config_path)
		} else if errors.Is(err, os.ErrPermission) {
			return nil, errors.New("Do not have permissions to read config file at: " + config_path)
		} else {
			return nil, errors.New("Unknown error for config file at: " + config_path + " -> " + err.Error())
		}
	}
	return config.ProcessDBConfigYAMLFile(config_path)
}
