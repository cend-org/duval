package configuration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	go_console "github.com/DrSmithFr/go-console"
)

const (
	RunnigModeTest  = 0
	RunningModeDev  = 1
	RunningModeProd = 2
)

type Config struct {
	Version                 string `toml:"version"`
	RunningMode             int    `toml:"version"`
	Port                    string `toml:"port"`
	Host                    string `toml:"host"`
	TokenSecret             string `toml:"token_secret"`
	DatabaseUserName        string `toml:"database_user_name"`
	DatabaseUserPassword    string `toml:"database_user_password"`
	DatabaseName            string `toml:"database_name"`
	DatabaseHost            string `toml:"database_host"`
	DatabasePort            string `toml:"database_port"`
	DatabaseConnexionString string
}

var App Config

func init() {
	_, err := toml.DecodeFile("config.toml", &App)
	if err != nil {
		App = Config{
			Version:                 "v0.0.1",
			Port:                    "8087",
			Host:                    "",
			RunningMode:             RunningModeProd,
			TokenSecret:             "437b059d-bd8b-40d5-920a-341bb8a3f15f",
			DatabaseUserName:        "root",
			DatabaseUserPassword:    "UnderAll4",
			DatabaseName:            "duval",
			DatabaseHost:            "db-duval.ctw4aeiceahd.eu-north-1.rds.amazonaws.com",
			DatabasePort:            "3306",
			DatabaseConnexionString: "",
		}
		fmt.Println("message: \"Your are using production configuration. If its not what you want, please configure your config.toml file.\"")
		cmd := go_console.NewScript().Build()
		cmd.PrintWarnings([]string{
			"Your are using production configuration. If its not what you want, please configure your config.toml file.",
		})

		err = nil
	}

	App.DatabaseConnexionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", App.DatabaseUserName,
		App.DatabaseUserPassword, App.DatabaseHost, App.DatabasePort,
		App.DatabaseName)
}

func (c *Config) IsDev() bool {
	return c.RunningMode == RunningModeDev
}

func (c *Config) IsProd() bool {
	return c.RunningMode == RunningModeProd
}

func (c *Config) IsTest() bool {
	return c.RunningMode == RunnigModeTest
}
