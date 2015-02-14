package database

import (
	"fmt"

	"git.ironlabs.com/greg/zenpager/config"
)

func GetDatasource() string {
	cfg := config.GetConfig()
	return fmt.Sprintf(
		"postgres://%s@%s/%s?sslmode=%s",
		cfg.DB_USER,
		cfg.DB_HOST,
		cfg.DB_NAME,
		cfg.DB_SSL,
	)
}
