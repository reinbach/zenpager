package models

import (
	"testing"

	"github.com/reinbach/zenpager/database"
)

// validate command group
func TestCommandGroupValidate(t *testing.T) {
	g := CommandGroup{}

	m := g.Validate()
	if len(m) == 0 {
		t.Errorf("Expected Command Group to fail validation")
	}

	g.Name = "G1"
	m = g.Validate()
	if len(m) != 0 {
		t.Errorf("Expected Command Group to pass validation")
	}
}

// create command group
func TestCommandGroupCreate(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G1"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected command group id to be set")
	}
}

// create command group with commands
func TestCommandGroupWithCommandsCreate(t *testing.T) {
	db := database.Connect()

	c := Command{Name: "CG2", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c.Create(db)

	g := CommandGroup{Name: "G5", Commands: []Command{c}}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	if g.ID == 0 {
		t.Errorf("Expected command group id to be set")
	}
}

// get all command groups
func TestCommandGroupGetAll(t *testing.T) {
	db := database.Connect()

	g1 := CommandGroup{Name: "G2"}

	r := g1.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	g2 := CommandGroup{
		Name: "G3",
	}

	r = g2.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	cg := CommandGroupGetAll(db)
	if len(cg) == 0 {
		t.Errorf("Expected to get command groups")
	}
}

// add command to group
func TestCommandGroupAddCommand(t *testing.T) {
	db := database.Connect()

	c := Command{Name: "CG1", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c.Create(db)

	g := CommandGroup{Name: "G4"}
	g.Create(db)

	r := g.AddCommand(db, &c)

	if r != true {
		t.Errorf("Expected success on adding command to group, got %v", r)
	}

	if len(g.Commands) != 1 {
		t.Errorf("Expected command in by group commands, got %v", g.Commands)
	}
}

// get command group
func TestCommandGroupGet(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G6"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	ng := CommandGroup{ID: 1234}
	ng.Get(db)

	if ng.Name == g.Name {
		t.Errorf("Expected to NOT get command group data")
	}

	ng.ID = g.ID
	ng.Get(db)

	if ng.Name != g.Name {
		t.Errorf("Expected to get command group data")
	}
}

// update command group
func TestCommandGroupUpdate(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G7"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	g.Name = "G8"
	r = g.Update(db)
	if r != true {
		t.Errorf("Expected success from command group update, got %v", r)
	}

	ng := CommandGroup{ID: g.ID}

	ng.Get(db)

	if ng.Name != "G8" {
		t.Errorf("Expected command group data to be updated, still %v",
			ng.Name,
		)
	}
}

// delete command group
func TestCommandGroupDelete(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G9"}

	r := g.Create(db)
	if r != true {
		t.Errorf("Expected success on command group create, got %v", r)
	}

	r = g.Delete(db)
	if r != true {
		t.Errorf("Expected success from command group delete, got %v", r)
	}

	ng := CommandGroup{ID: g.ID}
	ng.Get(db)

	if ng.Name != "" {
		t.Errorf("Expected command group to be deleted, still exists (%v)",
			ng.Name,
		)
	}
}

// get commands for command group
func TestCommandGroupCommands(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G10"}
	g.Create(db)

	g.GetCommands(db)

	if len(g.Commands) != 0 {
		t.Errorf("Expected no command group commands, got %v", g.Commands)
	}

	c1 := Command{Name: "CG3", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c1.Create(db)

	g.AddCommand(db, &c1)

	c2 := Command{Name: "CG4", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c2.Create(db)

	g.AddCommand(db, &c2)

	if len(g.Commands) != 2 {
		t.Errorf("Expected 2 command group commands, got %v", g.Commands)
	}
}

// remove command from command group
func TestCommandGroupRemoveCommand(t *testing.T) {
	db := database.Connect()

	g := CommandGroup{Name: "G11"}
	g.Create(db)

	c1 := Command{Name: "CG5", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c1.Create(db)

	g.AddCommand(db, &c1)

	c2 := Command{Name: "CG6", Command: "check_ssh $ARG1$ $HOSTADDRESS$"}
	c2.Create(db)

	g.AddCommand(db, &c2)

	if len(g.Commands) != 2 {
		t.Errorf("Expected 2 command group commands, got %v", len(g.Commands))
	}

	r := g.RemoveCommand(db, &c1)
	if r != true {
		t.Errorf("Expected success removing command from group, got %v", r)
	}

	g.GetCommands(db)

	if len(g.Commands) != 1 {
		t.Errorf("Expected 1 command in group, got %v", len(g.Commands))
	}
}
