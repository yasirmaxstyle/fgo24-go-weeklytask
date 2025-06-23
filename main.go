package main

import (
	"fmt"
	"go-cli/cmd"
	"log"
)

func main() {
	cli := cmd.NewCLI()

	fmt.Println("🚀 Starting App...")

	// Fetch menu data
	err := cli.FetchMenuData()
	if err != nil {
		log.Fatal("Failed to fetch menu data:", err)
	}

	cli.Run()
}
