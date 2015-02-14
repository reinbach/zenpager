package config

import (
	"fmt"
	"path/filepath"

	"github.com/pelletier/go-toml"

	"git.ironlabs.com/greg/zenpager/utils"
)

type Config struct {
	// database
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_SSL  string

	// session
	SESSION_HASHKEY string
	SESSION_SECRET  string
}

func GetConfig() Config {
	d := utils.GetAbsDir()
	config, err := toml.LoadFile(filepath.Join(d, "config.toml"))
	cfg := Config{}
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		// database
		cfg.DB_NAME = config.Get("postgresql.name").(string)
		cfg.DB_USER = config.Get("postgresql.user").(string)
		cfg.DB_PASS = config.Get("postgresql.host").(string)
		cfg.DB_HOST = config.Get("postgresql.password").(string)
		cfg.DB_SSL = config.Get("postgresql.sslmode").(string)

		// session
		cfg.SESSION_HASHKEY = config.Get("session.hashkey").(string)
		cfg.SESSION_SECRET = config.Get("session.secret").(string)
	}
	return cfg
}
