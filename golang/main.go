package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "create-password",
		Usage: "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello, World!")
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
