package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type Command struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Command string         `json:"command"`
	Groups  []CommandGroup `json:"groups"`
}

func CommandGetAll(db *sql.DB) []Command {
	commands := []Command{}
	rows, err := db.Query("SELECT id, name, command FROM command ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Commands not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var c Command
		err = rows.Scan(&c.ID, &c.Name, &c.Command)
		if err != nil {
			log.Println("Failed to get command data: ", err)
		}
		commands = append(commands, c)
	}

	return commands
}

func (c *Command) Create(db *sql.DB) bool {
	err := db.QueryRow(
		"INSERT INTO command (name, command) VALUES($1, $2) RETURNING id",
		c.Name,
		c.Command,
	).Scan(&c.ID)
	if err != nil {
		log.Printf("Failed to create command record. ", err)
		return false
	}

	return true
}

func (c *Command) Validate() []utils.Message {
	var errors []utils.Message
	if len(c.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	if len(c.Command) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Command is required."},
		)
	}
	return errors
}

func (c *Command) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name, command  FROM command WHERE id = $1",
		c.ID,
	).Scan(&c.Name, &c.Command)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Command not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (c *Command) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE command SET name = $1, command = $2 WHERE id = $3",
		c.Name, c.Command, c.ID)
	if err != nil {
		log.Printf("Failed to update command record. ", err)
		return false
	}
	return true
}

func (c *Command) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM command WHERE id = $1",
		c.ID)
	if err != nil {
		log.Printf("Failed to delete command record. ", err)
		return false
	}
	return true
}

func (c *Command) GetGroups(db *sql.DB) bool {
	rows, err := db.Query("SELECT g.id, g.name FROM command_group as g JOIN command_commandgroup AS sg ON g.id = sg.group_id WHERE sg.command_id = $1 ORDER BY g.name",
		c.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Command Group's Commands not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	c.Groups = []CommandGroup{}
	for rows.Next() {
		var g CommandGroup
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get command groups data: ", err)
		}
		c.Groups = append(c.Groups, g)
	}

	return true
}
