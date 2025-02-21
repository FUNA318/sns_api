package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ServerPort string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to load: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ServerPort: cfg.Section("api_info").Key("port").String(),
	}
}
