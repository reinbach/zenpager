package api

import (
	"testing"
)

// Routes
func TestRoutes(t *testing.T) {
	Routes("/api/v1/")
}

// Auth Routes
func TestAuthRoutes(t *testing.T) {
	AuthRoutes()
}

// UserRoutes
func TestUserRoutes(t *testing.T) {
	UserRoutes()
}

// Contact Routes
func TestContactRoutes(t *testing.T) {
	ContactRoutes()
}

// Server Routes
func TestServerRoutes(t *testing.T) {
	ServerRoutes()
}
