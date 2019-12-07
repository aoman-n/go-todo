package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	DbHost     string
	DbPassword string
	DbPort     int
	DbUser     string
	Port       string
}

var Config ConfigList

func init() {
	cnf, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v \n", err.Error())
		os.Exit(1)
	}

	Config = ConfigList{
		DbHost:     cnf.Section("db").Key("host").String(),
		DbPort:     cnf.Section("db").Key("port").MustInt(),
		DbUser:     cnf.Section("db").Key("user").String(),
		DbPassword: cnf.Section("db").Key("password").String(),
		Port:       cnf.Section("server").Key("port").String(),
	}
}
