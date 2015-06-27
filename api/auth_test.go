package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
)

func SetupWebContext() web.C {
	var ctx = context.Background()
	c := web.C{}
	db := database.Connect()
	ctx = database.NewContext(ctx, db)
	webctx.Set(&c, ctx)
	return c
}

// Routes
func TestRoutes(t *testing.T) {
	AuthRoutes()
}

// UserRoutes
func TestUserRoutes(t *testing.T) {
	UserRoutes()
}

// login no payload
func TestLoginNoPayload(t *testing.T) {
	body := url.Values{}
	body.Set("email", "test@example.com")
	body.Set("password", "123")
	b := bytes.NewBufferString(body.Encode())

	r, err := http.NewRequest("POST", "/login", b)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Login(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("400 expected, got %v instead", w.Code)
	}
}

// login, invalid params
func TestLoginFailed(t *testing.T) {
	u := models.User{
		Email:    "test@example.com",
		Password: "123",
	}
	j, _ := json.Marshal(u)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/login", b)
	r.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Login(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("400 expected, got %v instead", w.Code)
	}
}

// login, validation
func TestLoginValidation(t *testing.T) {
	u := models.User{
		Email: "test@example.com",
	}
	j, _ := json.Marshal(u)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/login", b)
	r.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Login(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("400 expected, got %v instead", w.Code)
	}
}

// login
func TestLoginSuccess(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test-valid@example.com",
		Password: "123",
	}
	u.Create(db)

	j, _ := json.Marshal(u)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/login", b)
	r.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Login(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}
}

// Logout
func TestLogout(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test-valid@example.com",
		Password: "123",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	b := bytes.NewBufferString("")

	r, err := http.NewRequest("GET", "/logout", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	Logout(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}

}

// User Partial Update
func TestUserPartialUpdate(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test321@example.com",
		Password: "123",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	u.Password = "321"

	j, _ := json.Marshal(u)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PATCH", fmt.Sprintf("/%d", u.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", u.ID)}

	UserPartialUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("202 expected, got %v instead", w.Code)
	}
}

// User Partial Update, validation
func TestUserPartialUpdateValidation(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test321@example.com",
		Password: "123",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	j, _ := json.Marshal(u)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PATCH", "/12", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()

	UserPartialUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("404 expected, got %v instead", w.Code)
	}
}

// User Partial Update, invalid
func TestUserPartialUpdateInvalid(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test321@example.com",
		Password: "123",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	uw := models.User{}
	j, _ := json.Marshal(uw)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PATCH", fmt.Sprintf("/%d", u.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", u.ID)}

	UserPartialUpdate(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%s expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}
