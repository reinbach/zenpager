package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
)

// Server List
func TestServerList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Server Item
func TestServerItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-0@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", s.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	ServerItem(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	// not found
	r, err = http.NewRequest("POST", "/321", nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ServerItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Server Add
func TestServerAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name string
		URL  string
	}
	d := Data{Name: "test", URL: "127.0.0.1"}
	j, _ := json.Marshal(d)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	// attempt to add again
	w = httptest.NewRecorder()
	ServerAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Server Add invalid data
func TestServerAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-2@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		ID int64
	}
	d := Data{ID: 321}
	j, _ := json.Marshal(d)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", "/", b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Server Update
func TestServerUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	s.Name = "Jane"
	j, _ := json.Marshal(s)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PUT", fmt.Sprintf("/%d", s.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	ServerUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	s.Get(db)
	if s.Name != "Jane" {
		t.Errorf("Expected data to be updated")
	}
}

// Server Delete
func TestServerDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", s.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	ServerDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Server Get Groups
func TestServerGetGroups(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-8@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	g1 := models.ServerGroup{Name: "CG5"}
	g1.Create(db)
	g1.AddServer(db, &s)

	g2 := models.ServerGroup{Name: "CG6"}
	g2.Create(db)
	g2.AddServer(db, &s)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/groups/", s.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	ServerGroups(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ServerGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}
