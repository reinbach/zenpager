package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type ContactGroup struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Contacts []Contact `json:"contacts"`
}

func ContactGroupGetAll(db *sql.DB) []ContactGroup {
	groups := []ContactGroup{}
	rows, err := db.Query("SELECT id, name FROM contact_group ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact Groups not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var g ContactGroup
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get contact group data: ", err)
		}
		groups = append(groups, g)
	}

	return groups
}

func (g *ContactGroup) Validate() []utils.Message {
	var errors []utils.Message
	if len(g.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	return errors
}

func (g *ContactGroup) Create(db *sql.DB) bool {
	err := db.QueryRow("INSERT INTO contact_group (name) VALUES($1) RETURNING id",
		g.Name).Scan(&g.ID)
	if err != nil {
		log.Printf("Failed to create contact group record. ", err)
		return false
	}
	log.Printf("Created contact group record.")

	return true
}

func (g *ContactGroup) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name FROM contact_group WHERE id = $1",
		g.ID,
	).Scan(&g.Name)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact Group not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (g *ContactGroup) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE contact_group SET name = $1 WHERE id = $2",
		g.Name, g.ID)
	if err != nil {
		log.Printf("Failed to update contact group record. ", err)
		return false
	}
	return true
}

func (g *ContactGroup) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM contact_group WHERE id = $1",
		g.ID)
	if err != nil {
		log.Printf("Failed to delete contact group record. ", err)
		return false
	}
	return true
}

func (g *ContactGroup) GetContacts(db *sql.DB) bool {
	rows, err := db.Query("SELECT c.id, c.name, u.email FROM contact_contact as c JOIN auth_user AS u on c.user_id = u.id JOIN contact_contactgroup AS cg ON c.id = cg.contact_id WHERE cg.group_id = $1 ORDER BY c.name",
		g.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact Group's Contacts not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	g.Contacts = []Contact{}
	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.ID, &c.Name, &c.User.Email)
		if err != nil {
			log.Println("Failed to get contact group contacts data: ", err)
		}
		g.Contacts = append(g.Contacts, c)
	}

	return true
}

func (g *ContactGroup) AddContact(db *sql.DB, c *Contact) bool {
	_, err := db.Exec("INSERT INTO contact_contactgroup (contact_id, group_id) VALUES($1, $2)",
		c.ID, g.ID)
	if err != nil {
		log.Printf("Failed to create contactgroup record. ", err)
		return false
	}
	log.Printf("Created contactgroup record.")

	g.Contacts = append(g.Contacts, *c)

	return true
}

func (g *ContactGroup) RemoveContact(db *sql.DB, c *Contact) bool {
	_, err := db.Exec("DELETE FROM contact_contactgroup WHERE contact_id = $1 AND group_id = $2",
		c.ID, g.ID)
	if err != nil {
		log.Printf("Failed to remove contact from group. ", err)
		return false
	}
	return true
}
