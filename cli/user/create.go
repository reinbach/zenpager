package user

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
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Print("Confirm Password: ")
	fmt.Scanln(&password_confirm)
}

func CreateUser() {
	// Create a user in the system
	fmt.Print("Email address: ")
	fmt.Scanln(&email)

	for password == "" || password != password_confirm {
		GetPassword()
	}

	user := auth.User{
		Email:    email,
		Password: database.Encrypt(password),
	}
	db := database.Connect()
	if user.Create(db) {
		fmt.Println("User was created.")
	} else {
		fmt.Println("There was an issue creating the user.")
	}
}
