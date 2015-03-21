package main

import (
	"fmt"
	"os"

	"github.com/reinbach/zenpager/cli/database"
	"github.com/reinbach/zenpager/cli/user"
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
