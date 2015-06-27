package models

import (
	"database/sql"
	"log"
)

type Token struct {
	ID    int64  `json:"id"`
	Token string `json:"token"`
	User  User   `json:"user"`
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
