package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
)

// User Partial Update
func TestUserPartialUpdate(t *testing.T) {
	db := database.Connect()

	u := models.User{
		Email:    "test1@example.com",
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
		Email:    "test2@example.com",
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
		Email:    "test3@example.com",
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
