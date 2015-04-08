package main

import (
	"fmt"
	"os"

	"github.com/reinbach/zenpager/cli/user"
	"github.com/reinbach/zenpager/database"
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
