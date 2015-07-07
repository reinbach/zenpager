package models

import (
	"net/url"
	"testing"

	"github.com/reinbach/zenpager/database"
)

// validate server
func TestServerValidate(t *testing.T) {
	s := Server{}

	m := s.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Server to fail validation")
	}

	s.Name = "Joe"
	m = s.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Server to fail validation")
	}

	s.URL = url.URL{Host: "127.0.0.1"}
	m = s.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Server to pass validation")
	}
}

// create server
func TestServerCreate(t *testing.T) {
	db := database.Connect()

	s := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}

	r := s.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}
}

// get server
func TestServerGet(t *testing.T) {
	db := database.Connect()

	s := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}

	r := s.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}

	ns := Server{
		ID: 1234,
	}
	ns.Get(db)

	if ns.Name == s.Name {
		t.Errorf("Expected to NOT get server data")
	}

	ns.ID = s.ID
	ns.Get(db)

	if ns.Name != s.Name || ns.URL != s.URL {
		t.Errorf("Expected to get server data")
	}
}

// update server
func TestServerUpdate(t *testing.T) {
	db := database.Connect()

	s := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}

	r := s.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}

	s.Name = "Jane"
	r = s.Update(db)
	if r != true {
		t.Errorf("Expected success from server update, got %v", r)
	}

	ns := Server{
		ID: s.ID,
	}

	ns.Get(db)

	if ns.Name != "Jane" {
		t.Errorf("Expected data to be updated, still %v", ns.Name)
	}
}

// delete server
func TestServerDelete(t *testing.T) {
	db := database.Connect()

	s := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}

	r := s.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}

	r = s.Delete(db)
	if r != true {
		t.Errorf("Expected success from server delete, got %v", r)
	}

	ns := Server{
		ID: s.ID,
	}
	ns.Get(db)

	if ns.Name != "" {
		t.Errorf("Expected data to be deleted, still exists (%v)", ns.Name)
	}
}

// get all servers
func TestServerGetAll(t *testing.T) {
	db := database.Connect()

	s1 := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}

	r := s1.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}

	s2 := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.2"},
	}

	r = s2.Create(db)
	if r != true {
		t.Errorf("Expected success on server create, got %v", r)
	}

	s := ServerGetAll(db)
	if len(s) == 0 {
		t.Errorf("Expected to get servers")
	}
}

// get groups for server
func TestServerGetGroups(t *testing.T) {
	db := database.Connect()

	s := Server{
		Name: "Joe",
		URL:  url.URL{Host: "127.0.0.1"},
	}
	s.Create(db)

	g1 := ServerGroup{Name: "CG12"}
	g1.Create(db)
	g1.AddServer(db, &s)

	g2 := ServerGroup{Name: "CG13"}
	g2.Create(db)
	g2.AddServer(db, &s)

	r := s.GetGroups(db)
	if r != true {
		t.Errorf("Expected success in getting groups for server, got %v", r)
	}

	if len(s.Groups) != 2 {
		t.Errorf("Expected 2 groups for server, got %v", len(s.Groups))
	}
}
