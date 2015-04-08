package auth

import (
	"os"
	"testing"

	"github.com/reinbach/zenpager/database"
)

func TestMain(m *testing.M) {
	// setup
	os.Setenv("TEST", "true")
	database.DropTables()
	database.InitDB()

	r := m.Run()

	// teardown
	database.DropTables()

	os.Exit(r)
}

// validate both email and password, none given
func TestUserValidateEmailPassword(t *testing.T) {
	u := User{}
	m := u.Validate(true)
	if len(m) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(m))
	}
}

// validate just email, none given
func TestUserValidateEmailOnly(t *testing.T) {
	u := User{}
	m := u.Validate(false)
	if len(m) != 1 {
		t.Errorf("Expected 1 error, got %d", len(m))
	}
}

// validate valid email
func TestUserValidateEmail(t *testing.T) {
	u := User{
		Email: "test@example.com",
	}
	m := u.Validate(false)
	if len(m) > 0 {
		t.Errorf("Expected no errors, got %d", len(m))
	}
}

// validate invalid email
func TestUserValidateEmailInvalid(t *testing.T) {
	u := User{
		Email: "testexample.com",
	}
	m := u.Validate(false)
	if len(m) != 1 {
		t.Errorf("Expected 1 error, got %d", len(m))
	}
}

// validate valid email, and password, not given
func TestUserValidatePassword(t *testing.T) {
	u := User{
		Email: "test@example.com",
	}
	m := u.Validate(true)
	if len(m) != 1 {
		t.Errorf("Expected 1 error, got %d", len(m))
	}
}

// validate valid email, and password
func TestUserValidatePasswordEmpty(t *testing.T) {
	u := User{
		Email:    "test@example.com",
		Password: "",
	}
	m := u.Validate(true)
	if len(m) != 1 {
		t.Errorf("Expected 1 error, got %d", len(m))
	}
}

// validate valid email, and password
func TestUserValidate(t *testing.T) {
	u := User{
		Email:    "test@example.com",
		Password: "123",
	}
	m := u.Validate(true)
	if len(m) > 0 {
		t.Errorf("Expected no errors, got %d", len(m))
	}
}

// create user
func TestCreateUser(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %s", r)
	}
}

// create user, unique
func TestCreateUserInvalid(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test1@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %s", r)
	}

	r = u.Create(db)
	if r != false {
		t.Errorf("Expected failed create, got %s", r)
	}
}
