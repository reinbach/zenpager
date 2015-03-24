package auth

import (
	"database/sql"
	"log"
	"net/mail"

	"github.com/reinbach/zenpager/database"
)

type User struct {
	ID       int64
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

func (u *User) Login(db *sql.DB) bool {
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

func (u *User) Validate() []string {
	var errors []string
	if len(u.Email) < 1 {
		errors = append(errors, "Email is required.")
	} else {
		if _, err := mail.ParseAddress(u.Email); err != nil {
			errors = append(errors, "A valid Email address is required.")
		}
	}
	if len(u.Password) < 1 {
		errors = append(errors, "Password is required.")
	}
	return errors
}
