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

// Command Group List
func TestCommandGroupList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroupList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Command Group Item
func TestCommandGroupItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-group-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.CommandGroup{Name: "ACG1"}
	g.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	CommandGroupItem(c, w, r)
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
	CommandGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Command Group Add
func TestCommandGroupAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-group-api-2@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	type Data struct {
		Name string
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
	CommandGroupAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	gl := models.CommandGroupGetAll(db)
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
	CommandGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Command Group Add invalid data
func TestCommandGroupAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-group-api-3@example.com",
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
	CommandGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Command Group Update
func TestCommandGroupUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-4@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.CommandGroup{Name: "ACG3"}
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
	CommandGroupUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	CommandGroupUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	g.Get(db)
	if g.Name != "ACG4" {
		t.Errorf("Expected data to be updated")
	}
}

// Command Group Delete
func TestCommandGroupDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.CommandGroup{Name: "ACG5"}
	g.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroupDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	CommandGroupDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Command Group Commands
func TestCommandGroupCommands(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "CGC1",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}
	s.Create(db)

	g := models.CommandGroup{Name: "ACG5"}
	g.Create(db)

	g.AddCommand(db, &s)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/commands/", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroupCommands(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	CommandGroupCommands(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	CommandGroupCommands(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Command Group Add Command
func TestCommandGroupAddCommand(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "CGC2",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}
	s.Create(db)

	g := models.CommandGroup{Name: "ACG6"}
	g.Create(db)

	j, _ := json.Marshal(s)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d/commands/", g.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	CommandGroupAddCommand(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	CommandGroupAddCommand(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	CommandGroupAddCommand(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	g.GetCommands(db)
	if len(g.Commands) != 1 {
		t.Errorf("Expected a single command to be in group, got %v",
			len(g.Commands))
	}
}

// Command Group Remove Command
func TestCommandGroupRemoveCommand(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "command-api-7@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Command{
		Name:    "CGC3",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}
	s.Create(db)

	g := models.CommandGroup{Name: "ACG7"}
	g.Create(db)

	g.AddCommand(db, &s)

	r, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/%d/commands/%d", g.ID, s.ID),
		nil,
	)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	c.URLParams = map[string]string{
		"id":  "321",
		"cid": fmt.Sprintf("%d", s.ID),
	}
	CommandGroupRemoveCommand(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{
		"id":  fmt.Sprintf("%d", g.ID),
		"cid": fmt.Sprintf("%d", s.ID),
	}
	CommandGroupRemoveCommand(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	g.GetCommands(db)
	if len(g.Commands) != 0 {
		t.Errorf("Expected command to NOT be in group, got %v",
			len(g.Commands))
	}
}
