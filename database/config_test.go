package database

import (
	"strings"
	"testing"
)

func TestGetDataSource(t *testing.T) {
	r := GetDatasource()
	if strings.Contains(r, "postgres://") == false {
		t.Errorf("Expected datasource string, got %v", r)
	}
}
