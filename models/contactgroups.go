package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type ContactGroup struct {
	Contact *Contact
	Group   *Group
}

func (cg *ContactGroup) Create(db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO contact_contactgroup (contact_id, group_id) VALUES($1, $2)",
		cg.Contact.ID, cg.Group.ID)
	if err != nil {
		log.Printf("Failed to create contactgroup record. ", err)
		return false
	}
	log.Printf("Created contactgroup record.")
	return true
}

type Group struct {
	ID       int64
	Name     string
	Contacts []*Contact
}

func ContactGroupGetAll(db *sql.DB) []Group {
	groups := []Group{}
	rows, err := db.Query("SELECT id, name FROM contact_group ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact Groups not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var g Group
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get contact group data: ", err)
		}
		groups = append(groups, g)
	}

	return groups
}

func (g *Group) Validate() []utils.Message {
	var errors []utils.Message
	if len(g.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	return errors
}

func (g *Group) Create(db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO contact_group (name) VALUES($1)",
		g.Name)
	if err != nil {
		log.Printf("Failed to create contact group record. ", err)
		return false
	}
	log.Printf("Created contact group record.")

	var cg ContactGroup
	for _, c := range g.Contacts {
		cg = ContactGroup{Contact: c, Group: g}
		cg.Create(db)
	}

	return true
}
