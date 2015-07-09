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

// Command List
func TestCommandList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Command Item
func TestCommandItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-0@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
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
	CommandItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	CommandItem(c, w, r)
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
	CommandItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Command Add
func TestCommandAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name    string
		Command string
	}
	d := Data{Name: "test", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
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
	CommandAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	// attempt to add again
	w = httptest.NewRecorder()
	CommandAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Command Add invalid data
func TestCommandAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-2@example.com",
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
	CommandAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Command Update
func TestCommandUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
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
	CommandUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	CommandUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	s.Get(db)
	if s.Name != "Jane" {
		t.Errorf("Expected data to be updated")
	}
}

// Command Delete
func TestCommandDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
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
	CommandDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	CommandDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Command Get Groups
func TestCommandGetGroups(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-8@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}
	s.Create(db)

	g1 := models.CommandGroup{Name: "CG5"}
	g1.Create(db)
	g1.AddCommand(db, &s)

	g2 := models.CommandGroup{Name: "CG6"}
	g2.Create(db)
	g2.AddCommand(db, &s)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/groups/", s.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", s.ID)}
	CommandGroups(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	CommandGroups(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}
