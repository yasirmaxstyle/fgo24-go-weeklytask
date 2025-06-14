package main

import (
	"fmt"
	"go-cli/utils"
	"os"
)

func main() {
	cli := utils.NewCLI()

	fmt.Println("🚀 Starting App...")

	// Fetch menu data
	if err := cli.FetchMenuData(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	cli.Run()
}
