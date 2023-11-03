package config

import (
	"encoding/json"
	"os"
)

type db struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	redis    string
}
type dbslave struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
}

type auth struct {
	Token string `json:"token"`
	Key   string `json:"key"`
}

type url struct {
	Jira       string `json:"jira"`
	Confluence string `json:"confluence"`
	Jenkins    string `json:"jenkins"`
}

type Configuration struct {
	Db      db      `json:"db"`
	Dbslave dbslave `json:"dbslave"`
	Auth    auth    `json:"auth"`
	URL     url     `json:"url"`
	Mode    int     `json:"mode"`
}

func Load(env string) Configuration {

	var configFile string

	switch env {
	case "dev":
		configFile = "/home/yohan/workspace/cmd/config/config_dev.json"
	case "prd":
		configFile = "/home/yohan/workspace/cmd/config/config_prd.json"
	case "test":
		configFile = "/home/yohan/workspace/cmd/config/config_test.json"
	default:
		panic("Input ENV")
	}

	file, err := os.Open(configFile)

	if err != nil {
		panic("Error opening configuration file: " + err.Error())
	}
	defer file.Close()

	// Parse the JSON configuration
	var config Configuration
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		panic("Error decoding configuration: " + err.Error())
	}

	return config
}
