package config

import (
	"log"

	"github.com/caioeverest/gocfg"
)

type Config struct {
	ENV      string
	HTTP     HTTPConfig
	Database DBConfig
}

type HTTPConfig struct {
	Port      int
	Greetings string
}

type DBConfig struct {
	Host     string
	Port     int
	DbName   string
	User     string
	Password string
}

var conf = &Config{}

const (
	DEV  = "development"
	PROD = "production"
)

// Load configuration on application.yml
func Start() {
	if err := gocfg.Load(conf, gocfg.YAML); err != nil {
		log.Panicf("ERROR reading config, %+v", err)
	}
}

func IsDevelopment() bool { return conf.ENV == DEV }

func Get() *Config { return conf }
