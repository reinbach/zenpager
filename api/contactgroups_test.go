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

// Contact Group List
func TestContactGroupList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Contact Group Item
func TestContactGroupItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-group-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.Group{Name: "ACG1"}
	g.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupItem(c, w, r)
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
	ContactGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Contact Group Add
func TestContactGroupAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-group-api-2@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{Name: "ACG2"}
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
	ContactGroupAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	gl := models.ContactGroupGetAll(db)
	f := false
	for _, g := range gl {
		if g.Name == d.Name {
			f = true
		}
	}
	if f != true {
		t.Errorf("Expected data to be saved")
	}

	// attempt to add again
	w = httptest.NewRecorder()
	ContactGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Contact Group Add invalid data
func TestContactGroupAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-group-api-3@example.com",
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
	ContactGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Contact Group Update
func TestContactGroupUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-4@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.Group{Name: "ACG3"}
	g.Create(db)

	g.Name = "ACG4"
	j, _ := json.Marshal(g)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PUT", fmt.Sprintf("/%d", g.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	g.Get(db)
	if g.Name != "ACG4" {
		t.Errorf("Expected data to be updated")
	}
}

// Contact Group Delete
func TestContactGroupDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.Group{Name: "ACG5"}
	g.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}
