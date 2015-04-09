package user

import (
	"fmt"

	"github.com/reinbach/zenpager/auth"
	"github.com/reinbach/zenpager/database"
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
		Password: password,
	}
	db := database.Connect()
	if user.Create(db) {
		fmt.Println("User was created.")
	} else {
		fmt.Println("There was an issue creating the user.")
	}
}
