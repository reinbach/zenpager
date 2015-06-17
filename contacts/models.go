package contacts

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/auth"
	"github.com/reinbach/zenpager/utils"
)

type Contact struct {
	ID     int64
	Name   string `json:"name"`
	User   auth.User
	Groups []*Group
}

func (c *Contact) Create(db *sql.DB) bool {
	_, err := db.Exec(
		"INSERT INTO contact_contact (name, user_id) VALUES($1, $2)",
		c.Name,
		c.User.ID,
	)
	if err != nil {
		log.Printf("Failed to create contact record. ", err)
		return false
	}
	log.Printf("Created contact record.")

	var cg ContactGroup
	for _, g := range c.Groups {
		cg = ContactGroup{Contact: c, Group: g}
		cg.Create(db)
	}

	return true
}

func (c *Contact) Validate() []utils.Message {
	var errors []utils.Message
	if len(c.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	return errors
}

func (c *Contact) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name, user_id FROM contact_contact WHERE id = $1",
		c.ID,
	).Scan(&c.Name, &c.User.ID)
	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (c *Contact) GetByUser(db *sql.DB) {
	err := db.QueryRow(
		"SELECT id, name FROM contact_contact WHERE user_id = $1",
		c.User.ID,
	).Scan(&c.ID, &c.Name)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("Contact does not exist.")
	case err != nil:
		log.Fatal(err)
	}
}

func (c *Contact) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE contact_contact SET name = $1 WHERE id = $2",
		c.Name, c.ID)
	if err != nil {
		log.Printf("Failed to update contact record. ", err)
		return false
	}
	return true
}

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
