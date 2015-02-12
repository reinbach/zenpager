package auth

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/database"
)

type User struct {
	ID       int64
	Email    string
	Password string
}

type UserDB interface {
	// GetUser retrieves a specific user from the
	// database for the given ID.
	GetUser(id int64) (*User, error)
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
	return database.Validate(u.Password, password)
}

// GetUser retrieves a specific user from the
// database for the given ID.
func GetUser(c web.C, id int64) (*User, error) {
	var db = database.FromContext(c)
	fmt.Println("db: ", db)
	var user = new(User)
	user.ID = 1
	user.Email = "test@example.com"
	return user, nil
}
