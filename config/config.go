package config

import (
	"fmt"
	"github.com/pelletier/go-toml"
)

type Config struct {
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_SSL  string
}

func GetConfig() Config {
	config, err := toml.LoadFile("config.toml")
	cfg := Config{}
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		cfg.DB_NAME = config.Get("postgresql.name").(string)
		cfg.DB_USER = config.Get("postgresql.user").(string)
		cfg.DB_PASS = config.Get("postgresql.host").(string)
		cfg.DB_HOST = config.Get("postgresql.password").(string)
		cfg.DB_SSL = config.Get("postgresql.sslmode").(string)
	}
	return cfg
}
