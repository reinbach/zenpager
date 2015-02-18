package auth

import (
	"database/sql"
	"log"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/database"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

func (u *User) Login(c web.C) bool {
	db := database.FromContext(c)
	var password string
	err := db.QueryRow("SELECT password FROM auth_user where email = $1",
		u.Email).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that Email.")
	case err != nil:
		log.Fatal(err)
	}
	return database.EncryptMatch(password, u.Password)
}

func (u *User) Create(db *sql.DB) bool {
	_, err := db.Exec("INSERT INTO auth_user (email, password) VALUES($1, $2)",
		u.Email, u.Password)
	if err != nil {
		log.Printf("Failed to create user record. ", err)
		return false
	}
	log.Printf("Created user record.")
	return true
}
