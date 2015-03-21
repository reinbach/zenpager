package config

import (
	"fmt"
	"path/filepath"

	"github.com/pelletier/go-toml"

	"github.com/reinbach/zenpager/utils"
)

type Config struct {
	// database
	DB_NAME string
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_SSL  string

	// session
	SESSION_HASH_KEY  string
	SESSION_BLOCK_KEY string
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
		cfg.SESSION_HASH_KEY = config.Get("session.hash_key").(string)
		cfg.SESSION_BLOCK_KEY = config.Get("session.block_key").(string)
	}
	return cfg
}
