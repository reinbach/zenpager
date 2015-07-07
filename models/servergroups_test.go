package models

import (
	"net/url"
	"testing"

	"github.com/reinbach/zenpager/database"
)

// validate server group
func TestServerGroupValidate(t *testing.T) {
	g := ServerGroup{}

	m := g.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Server Group to fail validation")
	}

	g.Name = "G1"
	m = g.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Server Group to pass validation")
	}
}

// create server group
func TestServerGroupCreate(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G1"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected server group id to be set")
	}
}

// create server group with servers
func TestServerGroupWithServersCreate(t *testing.T) {
	db := database.Connect()

	s := Server{Name: "CG2", URL: url.URL{Host: "127.0.0.1"}}
	s.Create(db)

	g := ServerGroup{Name: "G5", Servers: []Server{s}}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected server group id to be set")
	}
}

// get all server groups
func TestServerGroupGetAll(t *testing.T) {
	db := database.Connect()

	g1 := ServerGroup{Name: "G2"}

	r := g1.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	g2 := ServerGroup{
		Name: "G3",
	}

	r = g2.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	gs := ServerGroupGetAll(db)
	if len(gs) == 0 {
		t.Errorf("Expected to get server groups")
	}
}

// add server to group
func TestServerGroupAddServer(t *testing.T) {
	db := database.Connect()

	s := Server{Name: "CG1", URL: url.URL{Host: "127.0.0.1"}}
	s.Create(db)

	g := ServerGroup{Name: "G4"}
	g.Create(db)

	r := g.AddServer(db, &s)

	if r != true {
		t.Errorf("Expected success on adding server to group, got %v", r)
	}

	if len(g.Servers) != 1 {
		t.Errorf("Expected server in by group servers, got %v", g.Servers)
	}
}

// get server group
func TestServerGroupGet(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G6"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	ng := ServerGroup{ID: 1234}
	ng.Get(db)

	if ng.Name == g.Name {
		t.Errorf("Expected to NOT get server group data")
	}

	ng.ID = g.ID
	ng.Get(db)

	if ng.Name != g.Name {
		t.Errorf("Expected to get server group data")
	}
}

// update server group
func TestServerGroupUpdate(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G7"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	g.Name = "G8"
	r = g.Update(db)
	if r != true {
		t.Errorf("Expected success from server group update, got %v", r)
	}

	ng := ServerGroup{ID: g.ID}

	ng.Get(db)

	if ng.Name != "G8" {
		t.Errorf("Expected server group data to be updated, still %v",
			ng.Name,
		)
	}
}

// delete server group
func TestServerGroupDelete(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G9"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on server group create, got %v", r)
	}

	r = g.Delete(db)
	if r != true {
		t.Errorf("Expected success from server group delete, got %v", r)
	}

	ng := ServerGroup{ID: g.ID}
	ng.Get(db)

	if ng.Name != "" {
		t.Errorf("Expected server group to be deleted, still exists (%v)",
			ng.Name,
		)
	}
}

// get servers for server group
func TestServerGroupServers(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G10"}
	g.Create(db)

	g.GetServers(db)

	if len(g.Servers) != 0 {
		t.Errorf("Expected no server group servers, got %v", g.Servers)
	}

	s1 := Server{Name: "CG3", URL: url.URL{Host: "127.0.0.1"}}
	s1.Create(db)

	g.AddServer(db, &s1)

	s2 := Server{Name: "CG4", URL: url.URL{Host: "127.0.0.1"}}
	s2.Create(db)

	g.AddServer(db, &s2)

	if len(g.Servers) != 2 {
		t.Errorf("Expected 2 server group servers, got %v", g.Servers)
	}
}

// remove server from server group
func TestServerGroupRemoveServer(t *testing.T) {
	db := database.Connect()

	g := ServerGroup{Name: "G11"}
	g.Create(db)

	s1 := Server{Name: "CG5", URL: url.URL{Host: "127.0.0.1"}}
	s1.Create(db)

	g.AddServer(db, &s1)

	s2 := Server{Name: "CG6", URL: url.URL{Host: "127.0.0.1"}}
	s2.Create(db)

	g.AddServer(db, &s2)

	if len(g.Servers) != 2 {
		t.Errorf("Expected 2 server group servers, got %v", len(g.Servers))
	}

	r := g.RemoveServer(db, &s1)
	if r != true {
		t.Errorf("Expected success removing server from group, got %v", r)
	}

	g.GetServers(db)

	if len(g.Servers) != 1 {
		t.Errorf("Expected 1 server in group, got %v", len(g.Servers))
	}
}
