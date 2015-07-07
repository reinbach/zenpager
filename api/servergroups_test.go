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

// Server Group List
func TestServerGroupList(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroupList(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Server Group Item
func TestServerGroupItem(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-group-api-1@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.ServerGroup{Name: "ACG1"}
	g.Create(db)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ServerGroupItem(c, w, r)
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
	ServerGroupItem(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Server Group Add
func TestServerGroupAdd(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-group-api-2@example.com",
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
	ServerGroupAdd(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	gl := models.ServerGroupGetAll(db)
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
	ServerGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Server Group Add invalid data
func TestServerGroupAddInvalidData(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-group-api-3@example.com",
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
	ServerGroupAdd(c, w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("%v expected, got %v instead", http.StatusBadRequest, w.Code)
	}
}

// Server Group Update
func TestServerGroupUpdate(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-4@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.ServerGroup{Name: "ACG3"}
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
	ServerGroupUpdate(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ServerGroupUpdate(c, w, r)
	if w.Code != http.StatusAccepted {
		t.Errorf("%v expected, got %v instead", http.StatusAccepted, w.Code)
	}

	g.Get(db)
	if g.Name != "ACG4" {
		t.Errorf("Expected data to be updated")
	}
}

// Server Group Delete
func TestServerGroupDelete(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	g := models.ServerGroup{Name: "ACG5"}
	g.Create(db)

	r, err := http.NewRequest("DELETE", fmt.Sprintf("/%d", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroupDelete(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ServerGroupDelete(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}
}

// Server Group Servers
func TestServerGroupServers(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-5@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "CGC1",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	g := models.ServerGroup{Name: "ACG5"}
	g.Create(db)

	g.AddServer(db, &s)

	r, err := http.NewRequest("GET", fmt.Sprintf("/%d/servers/", g.ID), nil)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroupServers(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ServerGroupServers(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ServerGroupServers(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}
}

// Server Group Add Server
func TestServerGroupAddServer(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-6@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "CGC2",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	g := models.ServerGroup{Name: "ACG6"}
	g.Create(db)

	j, _ := json.Marshal(s)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("POST", fmt.Sprintf("/%d/servers/", g.ID), b)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("X-Access-Token", ut.Token)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	w := httptest.NewRecorder()
	c := SetupWebContext()
	ServerGroupAddServer(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": fmt.Sprintf("%d", g.ID)}
	ServerGroupAddServer(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{"id": "321"}
	ServerGroupAddServer(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	g.GetServers(db)
	if len(g.Servers) != 1 {
		t.Errorf("Expected a single server to be in group, got %v",
			len(g.Servers))
	}
}

// Server Group Remove Server
func TestServerGroupRemoveServer(t *testing.T) {
	db := database.Connect()
	u := models.User{
		Email: "server-api-7@example.com",
	}
	u.Create(db)
	ut, _ := u.AddToken(db)

	s := models.Server{
		Name: "CGC3",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	g := models.ServerGroup{Name: "ACG7"}
	g.Create(db)

	g.AddServer(db, &s)

	r, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/%d/servers/%d", g.ID, s.ID),
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
		"sid": fmt.Sprintf("%d", s.ID),
	}
	ServerGroupRemoveServer(c, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("%v expected, got %v instead", http.StatusNotFound, w.Code)
	}

	w = httptest.NewRecorder()
	c.URLParams = map[string]string{
		"id":  fmt.Sprintf("%d", g.ID),
		"sid": fmt.Sprintf("%d", s.ID),
	}
	ServerGroupRemoveServer(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("%v expected, got %v instead", http.StatusOK, w.Code)
	}

	g.GetServers(db)
	if len(g.Servers) != 0 {
		t.Errorf("Expected server to NOT be in group, got %v",
			len(g.Servers))
	}
}
