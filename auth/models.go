package auth

import (
	"fmt"

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
	err := db.QueryRow("SELECT * FROM auth_user where email = $1", user)
	defer rows.Close()
	if err != nil {
		fmt.Println("Db Query error: ", err)
	}
	fmt.Println("rows: ", rows)
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
