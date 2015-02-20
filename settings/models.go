package settings

import (
	"database/sql"
	"log"
)

type Contact struct {
	ID     int64
	Name   string
	Email  string
	Groups []*Group
}

func (c *Contact) Create(db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO contact_contact (name, email) VALUES($1, $2)",
		c.Name, c.Email)
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
