package auth

import (
	"fmt"

	"code.google.com/p/go.net/context"

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

// GetUser retrieves a specific user from the
// database for the given ID.
func GetUser(c context.Context, id int64) (*User, error) {
	var db = database.FromContext(c)
	fmt.Println("db: ", db)
	var user = new(User)
	user.ID = 1
	user.Email = "test@example.com"
	return user, nil
}
