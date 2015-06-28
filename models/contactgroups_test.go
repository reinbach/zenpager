package models

import (
	"fmt"
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

// get all contact groups
func TestContactGroupGetAll(t *testing.T) {
	db := database.Connect()

	g1 := Group{
		Name: "G2",
	}

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
	fmt.Println("Contact ID: ", cg.Contact.ID)
	fmt.Println("Group ID: ", cg.Group.ID)
	r := cg.Create(db)

	if r != true {
		t.Errorf("Exepected success on contact group relation, got %v", r)
	}
}
