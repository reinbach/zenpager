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

// Contact Group Contacts
func TestContactGroupContacts(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	ct := models.Contact{
		Name: "CGC1",
		User: u,
	}
	ct.Create(db)

	g := models.Group{Name: "ACG5"}
	g.Create(db)

	g.AddContact(db, &ct)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/contacts/", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupContacts(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupContacts(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ContactGroupContacts(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Contact Group Add Contact
func TestContactGroupAddContact(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	ct := models.Contact{
		Name: "CGC2",
		User: u,
	}
	ct.Create(db)

	g := models.Group{Name: "ACG6"}
	g.Create(db)

	j, _ := json.Marshal(ct)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d/contacts/", g.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupAddContact(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupAddContact(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ContactGroupAddContact(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	g.GetContacts(db)
	if len(g.Contacts) != 1 {
		t.Errorf("Expected a single contact to be in group, got %v",
			len(g.Contacts))
	}
}

// Contact Group Contact Delete
func TestContactGroupContactDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "contact-api-7@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	ct := models.Contact{
		Name: "CGC3",
		User: u,
	}
	ct.Create(db)

	g := models.Group{Name: "ACG7"}
	g.Create(db)

	g.AddContact(db, &ct)

	j, _ := json.Marshal(ct)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d/contacts/", g.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ContactGroupRemoveContact(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ContactGroupRemoveContact(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ContactGroupRemoveContact(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	g.GetContacts(db)
	if len(g.Contacts) != 0 {
		t.Errorf("Expected contact to NOT be in group, got %v",
			len(g.Contacts))
	}
}
