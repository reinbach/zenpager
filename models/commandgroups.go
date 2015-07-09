package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type CommandGroup struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

func CommandGroupGetAll(db *sql.DB) []CommandGroup {
	groups := []CommandGroup{}
	rows, err := db.Query("SELECT id, name FROM command_group ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Command Groups not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var g CommandGroup
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get command group data: ", err)
		}
		groups = append(groups, g)
	}

	return groups
}

func (g *CommandGroup) Validate() []utils.Message {
	var errors []utils.Message
	if len(g.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	return errors
}

func (g *CommandGroup) Create(db *sql.DB) bool {
	err := db.QueryRow("INSERT INTO command_group (name) VALUES($1) RETURNING id",
		g.Name).Scan(&g.ID)
	if err != nil {
		log.Printf("Failed to create command group record. ", err)
		return false
	}
	log.Printf("Created command group record.")

	return true
}

func (g *CommandGroup) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name FROM command_group WHERE id = $1",
		g.ID,
	).Scan(&g.Name)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Command Group not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (g *CommandGroup) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE command_group SET name = $1 WHERE id = $2",
		g.Name, g.ID)
	if err != nil {
		log.Printf("Failed to update command group record. ", err)
		return false
	}
	return true
}

func (g *CommandGroup) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM command_group WHERE id = $1",
		g.ID)
	if err != nil {
		log.Printf("Failed to delete command group record. ", err)
		return false
	}
	return true
}

func (g *CommandGroup) GetCommands(db *sql.DB) bool {
	rows, err := db.Query("SELECT c.id, c.name, c.command FROM command as c JOIN command_commandgroup AS cg ON c.id = cg.command_id WHERE cg.group_id = $1 ORDER BY c.name",
		g.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Command Group's Commands not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	g.Commands = []Command{}
	for rows.Next() {
		var c Command
		err = rows.Scan(&c.ID, &c.Name, &c.Command)
		if err != nil {
			log.Println("Failed to get command group commands data: ", err)
		}
		g.Commands = append(g.Commands, c)
	}

	return true
}

func (g *CommandGroup) AddCommand(db *sql.DB, c *Command) bool {
	_, err := db.Exec("INSERT INTO command_commandgroup (command_id, group_id) VALUES($1, $2)",
		c.ID, g.ID)
	if err != nil {
		log.Printf("Failed to create commandgroup record. ", err)
		return false
	}
	log.Printf("Created commandgroup record.")

	g.Commands = append(g.Commands, *c)

	return true
}

func (g *CommandGroup) RemoveCommand(db *sql.DB, c *Command) bool {
	_, err := db.Exec("DELETE FROM command_commandgroup WHERE command_id = $1 AND group_id = $2",
		c.ID, g.ID)
	if err != nil {
		log.Printf("Failed to remove command from group. ", err)
		return false
	}
	return true
}
