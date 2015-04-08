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
	err := db.QueryRow("SELECT id, password FROM auth_user WHERE email = $1",
		u.Email).Scan(&u.ID, &password)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that Email.")
	case err != nil:
		log.Fatal(err)
	}
	return database.EncryptMatch(password, u.Password)
}

func (u *User) Create(db *sql.DB) bool {
	_, err := db.Exec(
		"INSERT INTO auth_user (email, password) VALUES ($1, $2)",
		u.Email,
		u.Password,
	)
	if err != nil {
		log.Printf("Failed to create user record. ", err)
		return false
	}
	log.Printf("Created user record.")
	return true
}

func (u *User) Validate(validate_password bool) []Message {
	var errors []Message
	if len(u.Email) < 1 {
		errors = append(
			errors,
			Message{
				Type:    "danger",
				Content: "Email is required.",
			},
		)
	} else {
		if _, err := mail.ParseAddress(u.Email); err != nil {
			errors = append(
				errors,
				Message{
					Type:    "danger",
					Content: "A valid Email address is required.",
				},
			)
		}
	}
	if validate_password {
		if len(u.Password) < 1 {
			errors = append(
				errors,
				Message{
					Type:    "danger",
					Content: "Password is required.",
				},
			)
		}
	}
	return errors
}

func (u *User) Get(db *sql.DB) {
	err := db.QueryRow(
		"SELECT email FROM auth_user WHERE id = $1",
		u.ID,
	).Scan(&u.Email)

	switch {
	case err == sql.ErrNoRows:
		log.Println("User not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (u *User) Update(db *sql.DB) bool {
	if len(u.Password) > 0 {
		u.Password = database.Encrypt(u.Password)
		_, err := db.Exec(
			"UPDATE auth_user SET email = $1, password = $2 WHERE id = $3",
			u.Email, u.Password, u.ID,
		)
		if err != nil {
			log.Printf("Failed to update user record. ", err)
			return false
		}
		return true
	} else {
		_, err := db.Exec("UPDATE auth_user SET email = $1 WHERE id = $2",
			u.Email, u.ID)
		if err != nil {
			log.Printf("Failed to update user record. ", err)
			return false
		}
		return true
	}
}
