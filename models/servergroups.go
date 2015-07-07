package models

import (
	"database/sql"
	"log"

	"github.com/reinbach/zenpager/utils"
)

type ServerGroup struct {
	ID      int64    `json:"id"`
	Name    string   `json:"name"`
	Servers []Server `json:"servers"`
}

func ServerGroupGetAll(db *sql.DB) []ServerGroup {
	groups := []ServerGroup{}
	rows, err := db.Query("SELECT id, name FROM server_group ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Server Groups not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var g ServerGroup
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get server group data: ", err)
		}
		groups = append(groups, g)
	}

	return groups
}

func (g *ServerGroup) Validate() []utils.Message {
	var errors []utils.Message
	if len(g.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	return errors
}

func (g *ServerGroup) Create(db *sql.DB) bool {
	err := db.QueryRow("INSERT INTO server_group (name) VALUES($1) RETURNING id",
		g.Name).Scan(&g.ID)
	if err != nil {
		log.Printf("Failed to create server group record. ", err)
		return false
	}
	log.Printf("Created server group record.")

	return true
}

func (g *ServerGroup) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name FROM server_group WHERE id = $1",
		g.ID,
	).Scan(&g.Name)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Server Group not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (g *ServerGroup) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE server_group SET name = $1 WHERE id = $2",
		g.Name, g.ID)
	if err != nil {
		log.Printf("Failed to update server group record. ", err)
		return false
	}
	return true
}

func (g *ServerGroup) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM server_group WHERE id = $1",
		g.ID)
	if err != nil {
		log.Printf("Failed to delete server group record. ", err)
		return false
	}
	return true
}

func (g *ServerGroup) GetServers(db *sql.DB) bool {
	rows, err := db.Query("SELECT s.id, s.name, s.url FROM server as s JOIN server_servergroup AS sg ON s.id = sg.server_id WHERE sg.group_id = $1 ORDER BY s.name",
		g.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Server Group's Servers not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	g.Servers = []Server{}
	for rows.Next() {
		var s Server
		err = rows.Scan(&s.ID, &s.Name, &s.URL)
		if err != nil {
			log.Println("Failed to get server group servers data: ", err)
		}
		g.Servers = append(g.Servers, s)
	}

	return true
}

func (g *ServerGroup) AddServer(db *sql.DB, s *Server) bool {
	_, err := db.Exec("INSERT INTO server_servergroup (server_id, group_id) VALUES($1, $2)",
		s.ID, g.ID)
	if err != nil {
		log.Printf("Failed to create servergroup record. ", err)
		return false
	}
	log.Printf("Created servergroup record.")

	g.Servers = append(g.Servers, *s)

	return true
}

func (g *ServerGroup) RemoveServer(db *sql.DB, s *Server) bool {
	_, err := db.Exec("DELETE FROM server_servergroup WHERE server_id = $1 AND group_id = $2",
		s.ID, g.ID)
	if err != nil {
		log.Printf("Failed to remove server from group. ", err)
		return false
	}
	return true
}
