package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	NocreateCert bool   `json:"nocreatecert"`
	Ipv4         string `json:"ipv4"`
	Ipv6         string `json:"ipv6"`
}

func readConfig(filePath string) (Config, error) {
	var config Config
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(configFile, &config)
	return config, err
}
