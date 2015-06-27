package models

import (
	"os"
	"testing"

	"github.com/reinbach/zenpager/database"
)

func TestMain(m *testing.M) {
	// setup
	os.Setenv("TEST", "true")
	database.DropTables()
	database.InitDB()

	r := m.Run()

	// teardown
	database.DropTables()

	os.Exit(r)
}

// validate contact
func TestContactValidate(t *testing.T) {
	c := Contact{}

	m := c.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Contact to fail validation")
	}

	c.Name = "Joe"
	m = c.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Contact to pass validation")
	}
}

// create contact
func TestContactCreate(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "contacttest@example.com",
	}
	u.Create(db)

	c := Contact{
		Name: "Joe",
		User: u,
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}
}

// get contact
func TestContactGet(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "contacttest1@example.com",
	}
	u.Create(db)

	c := Contact{
		Name: "Joe",
		User: u,
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	nc := Contact{
		ID: 1234,
	}
	nc.Get(db)

	if nc.Name == c.Name {
		t.Errorf("Expected to NOT get contact data")
	}

	nc.ID = c.ID
	nc.Get(db)

	if nc.Name != c.Name {
		t.Errorf("Expected to get contact data")
	}

	if nc.User.ID != c.User.ID || nc.User.Email != c.User.Email {
		t.Errorf("Expected to get user data")
	}
}

// getbyuser contact
func TestContactGetByUser(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "contacttest2@example.com",
	}
	u.Create(db)

	c := Contact{
		Name: "Joe",
		User: u,
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	nc := Contact{}
	nc.GetByUser(db)

	if nc.Name == c.Name {
		t.Errorf("Expected to NOT get contact data")
	}

	nc.User.ID = u.ID
	nc.GetByUser(db)

	if nc.Name != c.Name {
		t.Errorf("Expected to get contact data")
	}
}

// update contact
func TestContactUpdate(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "contacttest3@example.com",
	}
	u.Create(db)

	c := Contact{
		Name: "Joe",
		User: u,
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	c.Name = "Jane"
	r = c.Update(db)
	if r != true {
		t.Errorf("Expected success from contact update, got %v", r)
	}

	nc := Contact{
		ID: c.ID,
	}

	nc.Get(db)

	if nc.Name != "Jane" {
		t.Errorf("Expected data to be updated, still %v", nc.Name)
	}
}

// delete contact
func TestContactDelete(t *testing.T) {
	db := database.Connect()
	u := User{
		Email: "contacttest4@example.com",
	}
	u.Create(db)

	c := Contact{
		Name: "Joe",
		User: u,
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	r = c.Delete(db)
	if r != true {
		t.Errorf("Expected success from contact delete, got %v", r)
	}

	nc := Contact{
		ID: c.ID,
	}
	nc.Get(db)

	if nc.Name != "" {
		t.Errorf("Expected data to be deleted, still exists (%v)", nc.Name)
	}
}

// get all contacts
func TestContactGetAll(t *testing.T) {
	db := database.Connect()

	u1 := User{
		Email: "contacttest5@example.com",
	}
	u1.Create(db)

	u2 := User{
		Email: "contacttest6@example.com",
	}
	u2.Create(db)

	c1 := Contact{
		Name: "Joe",
		User: u1,
	}

	r := c1.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	c2 := Contact{
		Name: "Joe",
		User: u2,
	}

	r = c2.Create(db)
	if r != true {
		t.Errorf("Expected success on contact create, got %v", r)
	}

	cs := ContactGetAll(db)
	if len(cs) == 0 {
		t.Errorf("Expected to get contacts")
	}
}
