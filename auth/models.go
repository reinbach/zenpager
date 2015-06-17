package auth

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"net/mail"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
	User  User   `json:"user"`
}

func RandomPassword() string {
	c := 20
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal("Failed to generate random password: ", err)
	}
	bytes.Equal(b, make([]byte, c))
	return fmt.Sprintf("%x", b)
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
	var p string
	if len(u.Password) < 1 {
		u.Password = RandomPassword()
	}
	p = database.Encrypt(u.Password)
	err := db.QueryRow(
		"INSERT INTO auth_user (email, password) VALUES ($1, $2) RETURNING id",
		u.Email,
		p,
	).Scan(&u.ID)
	if err != nil {
		log.Printf("Failed to create user record. ", err)
		return false
	}
	log.Printf("Created user record.")
	return true
}

func (u *User) Validate(validate_password bool) []utils.Message {
	var errors []utils.Message
	if len(u.Email) < 1 {
		errors = append(
			errors,
			utils.Message{
				Type:    "danger",
				Content: "Email is required.",
			},
		)
	} else {
		if _, err := mail.ParseAddress(u.Email); err != nil {
			errors = append(
				errors,
				utils.Message{
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
				utils.Message{
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

func (u *User) GetByEmail(db *sql.DB) {
	err := db.QueryRow(
		"SELECT id FROM auth_user WHERE email = $1",
		u.Email,
	).Scan(&u.ID)

	switch {
	case err == sql.ErrNoRows:
		log.Println("User not found.")
	case err != nil:
		log.Fatal(err)
	}
}

func (u *User) Update(db *sql.DB) bool {
	if u.ID == 0 {
		log.Printf("Invalid ID")
		return false
	}
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

func (u *User) AddToken(db *sql.DB) (Token, error) {
	var t = Token{
		Token: RandomPassword(),
		User:  *u,
	}

	err := db.QueryRow(
		"INSERT INTO auth_token (token, user_id) VALUES ($1, $2) RETURNING id",
		t.Token,
		u.ID,
	).Scan(&t.ID)
	if err != nil {
		log.Printf("Failed to create user token record. ", err)
		return t, err
	}

	return t, nil
}

func (t *Token) Get(db *sql.DB) bool {
	err := db.QueryRow("SELECT user_id FROM auth_token WHERE token = $1",
		t.Token).Scan(&t.User.ID)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No valid token.")
		return false
	case err != nil:
		log.Fatal(err)
		return false
	}
	return true
}

func RemoveToken(t string, db *sql.DB) {
	_, err := db.Exec(
		"DELETE FROM auth_token WHERE token = $1",
		t,
	)

	if err != nil {
		log.Println("Failed to remove token: ", err)
	}
}
