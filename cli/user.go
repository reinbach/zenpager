package cli

import (
	"fmt"

	"git.ironlabs.com/greg/zenpager/auth"
	"git.ironlabs.com/greg/zenpager/database"
)

var (
	email            string
	password         string
	password_confirm string
)

func GetPassword() {
	fmt.Println("Password: ")
	fmt.ScanIn(&password)
	fmt.Println("Confirm Password: ")
	fmt.ScanIn(&password_confirm)
}

func main() {
	// Create a user in the system
	fmt.Println("Email address: ")
	fmt.ScanIn(&email)

	for password == "" && password != password_confirm {
		GetPassword()
	}

	user := User{
		Email:    email,
		Password: password,
	}
	user.Create()
	fmt.Println("User has been created.")
}
