package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sifer169966/go-logger"
)

type Config struct {
	App      appConfig
	Postgres postgresConfig
}

type appConfig struct {
	HTTPPort string `envconfig:"APP_HTTP_PORT" example:"8443"`
	ENV      string `envconfig:"APP_ENV" example:"staging"`
	CodeName string `envconfig:"APP_CODE_NAME" example:"UNK"`
}

type postgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     string `envconfig:"POSTGRES_PORT"`
	User     string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
	DBName   string `envconfig:"POSTGRES_DB_NAME"`
	Debug    bool   `envconfig:"POSTGRES_DEBUG"`
}

var config Config

// Init is application config initialization ...
func Init() error {
	err := godotenv.Load()
	if err != nil {
		envFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if !envFileNotFound {
			logger.Info("cound not read environment", "error", err)
		} else {
			logger.Info("use OS environment")
		}
	}
	err = envconfig.Process("", &config)
	if err != nil {
		return err
	}
	return nil
}

// Get is use for export private variable which is config ...
func Get() *Config {
	return &config
}
