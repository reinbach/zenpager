package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type Contact struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	User   User    `json:"user"`
	Groups []Group `json:"groups"`
}

func ContactGetAll(db *sql.DB) []Contact {
	contacts := []Contact{}
	rows, err := db.Query("SELECT c.id, c.name, u.email FROM contact_contact as c JOIN auth_user AS u on c.user_id = u.id ORDER BY c.name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contacts not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var c Contact
		err = rows.Scan(&c.ID, &c.Name, &c.User.Email)
		if err != nil {
			log.Println("Failed to get contact data: ", err)
		}
		contacts = append(contacts, c)
	}

	return contacts
}

func (c *Contact) Create(db *sql.DB) bool {
	err := db.QueryRow(
		"INSERT INTO contact_contact (name, user_id) VALUES($1, $2) RETURNING id",
		c.Name,
		c.User.ID,
	).Scan(&c.ID)
	if err != nil {
		log.Printf("Failed to create contact record. ", err)
		return false
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
		"SELECT c.name, c.user_id, u.email  FROM contact_contact AS c JOIN auth_user AS u ON c.user_id = u.id WHERE c.id = $1",
		c.ID,
	).Scan(&c.Name, &c.User.ID, &c.User.Email)

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

func (c *Contact) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM contact_contact WHERE id = $1",
		c.ID)
	if err != nil {
		log.Printf("Failed to delete contact record. ", err)
		return false
	}
	return true
}

func (c *Contact) GetGroups(db *sql.DB) bool {
	rows, err := db.Query("SELECT g.id, g.name FROM contact_group as g JOIN contact_contactgroup AS cg ON g.id = cg.group_id WHERE cg.contact_id = $1 ORDER BY g.name",
		c.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Contact Group's Contacts not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	c.Groups = []Group{}
	for rows.Next() {
		var g Group
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get contact groups data: ", err)
		}
		c.Groups = append(c.Groups, g)
	}

	return true
}
