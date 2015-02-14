package main

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
	fmt.Scanln(&password)
	fmt.Println("Confirm Password: ")
	fmt.Scanln(&password_confirm)
}

func main() {
	// Create a user in the system
	fmt.Println("Email address: ")
	fmt.Scanln(&email)

	for password == "" && password != password_confirm {
		GetPassword()
	}

	user := auth.User{
		Email:    email,
		Password: database.Encrypt(password),
	}
	db := database.Connect(datasource)
	user.Create(db)
	fmt.Println("User has been created.")
}
