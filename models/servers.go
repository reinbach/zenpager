package models

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/reinbach/zenpager/utils"
)

type Server struct {
	ID     int64         `json:"id"`
	Name   string        `json:"name"`
	URL    url.URL       `json:"url"`
	Groups []ServerGroup `json:"groups"`
}

func ServerGetAll(db *sql.DB) []Server {
	servers := []Server{}
	rows, err := db.Query("SELECT id, name, url FROM server ORDER BY name")

	switch {
	case err == sql.ErrNoRows:
		log.Println("Servers not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var s Server
		err = rows.Scan(&s.ID, &s.Name, &s.URL.Host)
		if err != nil {
			log.Println("Failed to get server data: ", err)
		}
		servers = append(servers, s)
	}

	return servers
}

func (s *Server) Create(db *sql.DB) bool {
	err := db.QueryRow(
		"INSERT INTO server (name, url) VALUES($1, $2) RETURNING id",
		s.Name,
		s.URL.Host,
	).Scan(&s.ID)
	if err != nil {
		log.Printf("Failed to create server record. ", err)
		return false
	}

	return true
}

func (s *Server) Validate() []utils.Message {
	var errors []utils.Message
	if len(s.Name) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "Name is required."},
		)
	}
	if len(s.URL.Host) < 1 {
		errors = append(
			errors,
			utils.Message{Type: "danger", Content: "URL is required."},
		)
	}
	return errors
}

func (s *Server) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT name, url  FROM server WHERE id = $1",
		s.ID,
	).Scan(&s.Name, &s.URL.Host)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Server not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (s *Server) Update(db *sql.DB) bool {
	_, err := db.Exec("UPDATE server SET name = $1, url = $2 WHERE id = $3",
		s.Name, s.URL.Host, s.ID)
	if err != nil {
		log.Printf("Failed to update server record. ", err)
		return false
	}
	return true
}

func (s *Server) Delete(db *sql.DB) bool {
	_, err := db.Exec("DELETE FROM server WHERE id = $1",
		s.ID)
	if err != nil {
		log.Printf("Failed to delete server record. ", err)
		return false
	}
	return true
}

func (s *Server) GetGroups(db *sql.DB) bool {
	rows, err := db.Query("SELECT g.id, g.name FROM server_group as g JOIN server_servergroup AS sg ON g.id = sg.group_id WHERE sg.server_id = $1 ORDER BY g.name",
		s.ID,
	)

	switch {
	case err == sql.ErrNoRows:
		log.Println("Server Group's Servers not found.")
	case err != nil:
		log.Fatal(err)
	}

	defer rows.Close()
	s.Groups = []ServerGroup{}
	for rows.Next() {
		var g ServerGroup
		err = rows.Scan(&g.ID, &g.Name)
		if err != nil {
			log.Println("Failed to get server groups data: ", err)
		}
		s.Groups = append(s.Groups, g)
	}

	return true
}
