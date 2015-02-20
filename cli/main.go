package main

import (
	"fmt"
	"os"

	"git.ironlabs.com/greg/zenpager/cli/database"
	"git.ironlabs.com/greg/zenpager/cli/user"
)

func main() {
	var command = os.Args[1]

	switch command {
	case "migrate":
		database.Migrate()
	case "createuser":
		user.CreateUser()
	default:
		fmt.Println("Unknown command.")
	}
}
