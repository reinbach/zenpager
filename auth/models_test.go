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
		t.Errorf("Expected successful create, got %t", r)
	}
	if u.ID == 0 {
		t.Errorf("Expected ID to be set, got %d", u.ID)
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
		t.Errorf("Expected successful create, got %t", r)
	}

	r = u.Create(db)
	if r != false {
		t.Errorf("Expected failed create, got %t", r)
	}
}

// get user, empty
func TestGetUserEmpty(t *testing.T) {
	db := database.Connect()
	u := User{}
	u.Get(db)
	if u.Email != "" {
		t.Errorf("Expected not found, got %s", u.Email)
	}
}

// get user, empty
func TestGetUser(t *testing.T) {
	db := database.Connect()
	u1 := User{
		Email:    "test2@example.com",
		Password: "123",
	}
	r := u1.Create(db)
	if r != true {
		t.Errorf("Expected succesfull create, got %t", r)
	}

	u2 := User{
		ID: u1.ID,
	}
	u2.Get(db)
	if u2.Email != u1.Email {
		t.Errorf("Expected matching record, got %s", u2.Email)
	}
}

// update user, no password
func TestUpdateUserNoPassword(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "test3@example.com",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}
	u.Password = ""
	u.Email = "change@example.com"
	r = u.Update(db)
	if r != true {
		t.Errorf("Expected successful update, got %t", r)
	}
}

// update user, not valid
func TestUpdateUserNotValid(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "test4@example.com",
	}
	r := u.Update(db)
	if r != false {
		t.Errorf("Expected failed update, got %t", r)
	}
}

// update user, with password
func TestUpdateUserWithPassword(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test5@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	u.Password = "321"
	r = u.Update(db)
	if r != true {
		t.Errorf("Expected successful update, got %t", r)
	}
}

// get user, with email
func TestGetByEmail(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test5.5@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	nu := User{
		Email: "test5.5@example.com",
	}
	nu.GetByEmail(db)
	if nu.ID == 0 {
		t.Errorf("Expected successful get by email, got nothing")
	}
}

// get user, with email, does not exist
func TestGetByEmailInvalid(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "wrong-email@example.com",
	}
	u.GetByEmail(db)
	if u.ID != 0 {
		t.Errorf("Expected no result, got %v", u)
	}
}

// login valid
func TestLogin(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test6@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	l := u.Login(db)
	if l != true {
		t.Errorf("Expected successful login, got %t", l)
	}
}

// login invalid password
func TestLoginInValidPassword(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test7@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	u.Password = "321"
	l := u.Login(db)
	if l != false {
		t.Errorf("Expected invalid login, got %t", l)
	}
}

// login invalid email
func TestLoginInValidEmail(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test8@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	u.Email = "wrong@example.com"
	l := u.Login(db)
	if l != false {
		t.Errorf("Expected invalid login, got %t", l)
	}
}

// add token
func TestAddToken(t *testing.T) {
	db := database.Connect()
	u := User{
		Email:    "test9@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	rt, err := u.AddToken(db)
	if err != nil {
		t.Errorf("Expected to add token, got %t", err)
	}

	if rt.User.ID != u.ID {
		t.Errorf("Expected user on token to match user")
	}
}

// get token
func TestGetToken(t *testing.T) {
	db := database.Connect()

	u := User{
		Email:    "test10@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	ut, err := u.AddToken(db)
	if err != nil {
		t.Errorf("Expected token to be added, got %t", err)
	}

	b := ut.Get(db)
	if b != true {
		t.Errorf("Expected get token to return true")
	}
}

// get token, invalid
func TestGetTokenInvalid(t *testing.T) {
	db := database.Connect()

	u := User{
		Email:    "test11@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	_, err := u.AddToken(db)
	if err != nil {
		t.Errorf("Expected token to be added, got %t", err)
	}

	uwt := Token{
		Token: "321",
	}
	b := uwt.Get(db)
	if b != false {
		t.Errorf("Expected get token to return false")
	}
}

// delete token
func TestDeleteToken(t *testing.T) {
	db := database.Connect()

	u := User{
		Email:    "test12@example.com",
		Password: "123",
	}
	r := u.Create(db)
	if r != true {
		t.Errorf("Expected successful create, got %t", r)
	}

	ut, err := u.AddToken(db)
	if err != nil {
		t.Errorf("Expected token to be added, got %t", err)
	}

	RemoveToken(ut.Token, db)

	rt := ut.Get(db)
	if rt != false {
		t.Errorf("Expected get token to return false")
	}
}

// delete token, invalid
func TestDeleteTokenInvalid(t *testing.T) {
	db := database.Connect()

	ut := Token{
		Token: "321",
	}

	RemoveToken(ut.Token, db)
}
