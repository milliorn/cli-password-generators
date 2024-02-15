package main

import (
	"fmt"  // Importing the 'fmt' package for formatted I/O operations
	"log"  // Importing the 'log' package for logging errors
	"os"   // Importing the 'os' package for accessing command-line arguments and environment variables

	"github.com/urfave/cli/v2"  // Importing the 'github.com/urfave/cli/v2' package for building CLI applications
)

func main() {
	// Creating a new CLI application instance with name and usage description
	app := &cli.App{
		Name:  "create-password",  // Setting the name of the CLI application
		Usage: "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.",  // Setting the usage description of the CLI application
		Action: func(c *cli.Context) error {  // Defining the action to be executed when the CLI application is run
			fmt.Println("Hello, World!")  // Printing "Hello, World!" to the console
			return nil  // Returning nil to indicate successful execution of the action
		},
	}

	// Running the CLI application with the command-line arguments and handling any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)  // Logging fatal errors
	}
}
