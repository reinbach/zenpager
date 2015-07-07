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

// Contact List
func TestContactList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Contact Item
func TestContactItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-10@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := models.Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", cu.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	ContactItem(c, w, r)
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
	ContactItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Contact Add
func TestContactAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{Name: "test", Email: u.Email}
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
	ContactAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	cu := models.Contact{
		User: u,
	}
	cu.GetByUser(db)
	if cu.Name != d.Name {
		t.Errorf("Expected data to be saved")
	}

	// attempt to add again
	w = httptest.NewRecorder()
	ContactAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Contact Add invalid data
func TestContactAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-2@example.com",
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
	ContactAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Contact Add that creates user
func TestContactAddCreateUser(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-3@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{Name: "test", Email: "contact-api-4@example.com"}
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
	ContactAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}
}

// Contact Update
func TestContactUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := models.Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	cu.Name = "Jane"
	j, _ := json.Marshal(cu)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("PUT", fmt.Sprintf("/%d", cu.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	ContactUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	cu.Get(db)
	if cu.Name != "Jane" {
		t.Errorf("Expected data to be updated")
	}
}

// Contact Delete
func TestContactDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	cu := models.Contact{
		Name: "Joe",
		User: u,
	}
	cu.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", cu.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", cu.ID)}
	ContactDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Contact Get Groups
func TestContactGetGroups(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-8@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	ct := models.Contact{
		Name: "Joe",
		User: u,
	}
	ct.Create(db)

	g1 := models.ContactGroup{Name: "CG5"}
	g1.Create(db)
	g1.AddContact(db, &ct)

	g2 := models.ContactGroup{Name: "CG6"}
	g2.Create(db)
	g2.AddContact(db, &ct)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/groups/", ct.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", ct.ID)}
	ContactGroups(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ContactGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}
