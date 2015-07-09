package models

import (
	"testing"

	"github.com/reinbach/zenpager/database"
)

// validate command
func TestCommandValidate(t *testing.T) {
	c := Command{}

	m := c.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Command to fail validation")
	}

	c.Name = "Joe"
	m = c.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Command to fail validation")
	}

	c.Command = "check_ssh $ARG1$ $HOSTADDRESS$"
	m = c.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Command to pass validation")
	}
}

// create command
func TestCommandCreate(t *testing.T) {
	db := database.Connect()

	c := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}
}

// get command
func TestCommandGet(t *testing.T) {
	db := database.Connect()

	c := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}

	nc := Command{
		ID: 1234,
	}
	nc.Get(db)

	if nc.Name == c.Name {
		t.Errorf("Expected to NOT get command data")
	}

	nc.ID = c.ID
	nc.Get(db)

	if nc.Name != c.Name || nc.Command != c.Command {
		t.Errorf("Expected to get command data")
	}
}

// update command
func TestCommandUpdate(t *testing.T) {
	db := database.Connect()

	c := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}

	c.Name = "Jane"
	r = c.Update(db)
	if r != true {
		t.Errorf("Expected success from command update, got %v", r)
	}

	nc := Command{
		ID: c.ID,
	}

	nc.Get(db)

	if nc.Name != "Jane" {
		t.Errorf("Expected data to be updated, still %v", nc.Name)
	}
}

// delete command
func TestCommandDelete(t *testing.T) {
	db := database.Connect()

	c := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r := c.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}

	r = c.Delete(db)
	if r != true {
		t.Errorf("Expected success from command delete, got %v", r)
	}

	nc := Command{
		ID: c.ID,
	}
	nc.Get(db)

	if nc.Name != "" {
		t.Errorf("Expected data to be deleted, still exists (%v)", nc.Name)
	}
}

// get all commands
func TestCommandGetAll(t *testing.T) {
	db := database.Connect()

	c1 := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r := c1.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}

	c2 := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}

	r = c2.Create(db)
	if r != true {
		t.Errorf("Expected success on command create, got %v", r)
	}

	c := CommandGetAll(db)
	if len(c) == 0 {
		t.Errorf("Expected to get commands")
	}
}

// get groups for command
func TestCommandGetGroups(t *testing.T) {
	db := database.Connect()

	c := Command{
		Name:    "Joe",
		Command: "check_ssh $ARG1$ $HOSTADDRESS$",
	}
	c.Create(db)

	g1 := CommandGroup{Name: "CG12"}
	g1.Create(db)
	g1.AddCommand(db, &c)

	g2 := CommandGroup{Name: "CG13"}
	g2.Create(db)
	g2.AddCommand(db, &c)

	r := c.GetGroups(db)
	if r != true {
		t.Errorf("Expected success in getting groups for command, got %v", r)
	}

	if len(c.Groups) != 2 {
		t.Errorf("Expected 2 groups for command, got %v", len(c.Groups))
	}
}
