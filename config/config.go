package config

import (
	"golang_udemy/to-do-app/utils"
	"log"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigLIst struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigLIst

func init() {
	LoadConfig()
	utils.LoggingSetting(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)

	}
	Config = ConfigLIst{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
