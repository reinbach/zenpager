package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg := GetConfig()

	if cfg.DB_NAME == "" {
		t.Errorf("Expected DB_NAME, got nothing")
	}

	if cfg.SESSION_HASH_KEY == "" {
		t.Errorf("Expected SESSION_HASH_KEY, got nothing")
	}
}
