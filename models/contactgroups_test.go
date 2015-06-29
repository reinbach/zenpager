package models

import (
	"testing"

	"github.com/reinbach/zenpager/database"
)

// validate contact group
func TestContactGroupValidate(t *testing.T) {
	g := Group{}

	m := g.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Contact Group to fail validation")
	}

	g.Name = "G1"
	m = g.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Contact Group to pass validation")
	}
}

// create contact group
func TestContactGroupCreate(t *testing.T) {
	db := database.Connect()

	g := Group{Name: "G1"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected contact group id to be set")
	}
}

// create contact group with contacts
func TestContactGroupWithContactsCreate(t *testing.T) {
	db := database.Connect()

	u := User{Email: "cg2@example.com"}
	u.Create(db)

	c := Contact{Name: "CG2", User: u}
	c.Create(db)

	g := Group{Name: "G5", Contacts: []Contact{c}}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected contact group id to be set")
	}
}

// get all contact groups
func TestContactGroupGetAll(t *testing.T) {
	db := database.Connect()

	g1 := Group{Name: "G2"}

	r := g1.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	g2 := Group{
		Name: "G3",
	}

	r = g2.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	gs := ContactGroupGetAll(db)
	if len(gs) == 0 {
		t.Errorf("Expected to get contact groups")
	}
}

// create contact group relation
func TestContactGroupLinkCreate(t *testing.T) {
	db := database.Connect()

	u := User{Email: "cg1@example.com"}
	u.Create(db)

	c := Contact{Name: "CG1", User: u}
	c.Create(db)

	g := Group{Name: "G4"}
	g.Create(db)

	cg := ContactGroup{Contact: &c, Group: &g}
	r := cg.Create(db)

	if r != true {
		t.Errorf("Exepected success on contact group relation, got %v", r)
	}
}

// get contact group
func TestContactGroupGet(t *testing.T) {
	db := database.Connect()

	g := Group{Name: "G6"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	ng := Group{ID: 1234}
	ng.Get(db)

	if ng.Name == g.Name {
		t.Errorf("Expected to NOT get contact group data")
	}

	ng.ID = g.ID
	ng.Get(db)

	if ng.Name != g.Name {
		t.Errorf("Expected to get contact group data")
	}
}

// update contact group
func TestContactGroupUpdate(t *testing.T) {
	db := database.Connect()

	g := Group{Name: "G7"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	g.Name = "G8"
	r = g.Update(db)
	if r != true {
		t.Errorf("Expected success from contact group update, got %v", r)
	}

	ng := Group{ID: g.ID}

	ng.Get(db)

	if ng.Name != "G8" {
		t.Errorf("Expected contact group data to be updated, still %v",
			ng.Name,
		)
	}
}

// delete contact group
func TestContactGroupDelete(t *testing.T) {
	db := database.Connect()

	g := Group{Name: "G9"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on contact group create, got %v", r)
	}

	r = g.Delete(db)
	if r != true {
		t.Errorf("Expected success from contact group delete, got %v", r)
	}

	ng := Group{ID: g.ID}
	ng.Get(db)

	if ng.Name != "" {
		t.Errorf("Expected contact group to be deleted, still exists (%v)",
			ng.Name,
		)
	}
}

// get contacts for contact group
func TestContactGroupContacts(t *testing.T) {
	db := database.Connect()

	g := Group{Name: "G10"}
	g.Create(db)

	g.GetContacts(db)

	if len(g.Contacts) != 0 {
		t.Errorf("Expected no contact group contacts, got %v", g.Contacts)
	}

	u1 := User{Email: "cg3@example.com"}
	u1.Create(db)

	c1 := Contact{Name: "CG3", User: u1}
	c1.Create(db)

	cg1 := ContactGroup{Contact: &c1, Group: &g}
	cg1.Create(db)

	u2 := User{Email: "cg4@example.com"}
	u2.Create(db)

	c2 := Contact{Name: "CG4", User: u2}
	c2.Create(db)

	cg2 := ContactGroup{Contact: &c2, Group: &g}
	cg2.Create(db)

	r := g.GetContacts(db)
	if r != true {
		t.Errorf("Expected success from contact group contacts, got %v", r)
	}

	if len(g.Contacts) != 2 {
		t.Errorf("Expected 2 contact group contacts, got %v", g.Contacts)
	}
}
