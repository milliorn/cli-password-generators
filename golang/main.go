package main

import (
	"fmt" // Importing the 'fmt' package for formatted I/O operations
	"log" // Importing the 'log' package for logging errors
	"os"  // Importing the 'os' package for accessing command-line arguments and environment variables
	"sort"

	"github.com/urfave/cli/v2" // Importing the 'github.com/urfave/cli/v2' package for building CLI applications
)

func main() {
	// Creating a new CLI application instance with name and usage description
	app := &cli.App{
		Name:  "create-password",                                                                                                      // Setting the name of the CLI application
		Usage: "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.", // Setting the usage description of the CLI application
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "length",                 // Setting the name of the flag
				Aliases: []string{"l"},            // Setting the aliases of the flag
				Value:   "8",                      // Setting the default value of the flag
				Usage:   "Length of the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-numbers",                        // Setting the name of the flag
				Aliases: []string{"nn"},                      // Setting the aliases of the flag
				Value:   false,                               // Setting the default value of the flag
				Usage:   "Exclude numbers from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-symbols",                                // Setting the name of the flag
				Aliases: []string{"ns"},                              // Setting the aliases of the flag
				Value:   false,                                       // Setting the default value of the flag
				Usage:   "Exclude special symbols from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-uppercase",                                // Setting the name of the flag
				Aliases: []string{"nu"},                                // Setting the aliases of the flag
				Value:   false,                                         // Setting the default value of the flag
				Usage:   "Exclude uppercase letters from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-lowercase",                                // Setting the name of the flag
				Aliases: []string{"nl"},                                // Setting the aliases of the flag
				Value:   false,                                         // Setting the default value of the flag
				Usage:   "Exclude lowercase letters from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-ambiguous",                                   // Setting the name of the flag
				Aliases: []string{"na"},                                   // Setting the aliases of the flag
				Value:   false,                                            // Setting the default value of the flag
				Usage:   "Exclude ambiguous characters from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-similar",                                   // Setting the name of the flag
				Aliases: []string{"nsi"},                                // Setting the aliases of the flag
				Value:   false,                                          // Setting the default value of the flag
				Usage:   "Exclude similar characters from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-vowels",                        // Setting the name of the flag
				Aliases: []string{"nv"},                     // Setting the aliases of the flag
				Value:   false,                              // Setting the default value of the flag
				Usage:   "Exclude vowels from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-consonants",                        // Setting the name of the flag
				Aliases: []string{"nc"},                         // Setting the aliases of the flag
				Value:   false,                                  // Setting the default value of the flag
				Usage:   "Exclude consonants from the password", // Setting the usage description of the flag
			},

			&cli.BoolFlag{
				Name:    "no-letters",                        // Setting the name of the flag
				Aliases: []string{"nle"},                     // Setting the aliases of the flag
				Value:   false,                               // Setting the default value of the flag
				Usage:   "Exclude letters from the password", // Setting the usage description of the flag
			},
		},
		Action: func(c *cli.Context) error { // Defining the action to be executed when the CLI application is run
			fmt.Println("Hello, World!") // Printing "Hello, World!" to the console
			return nil                   // Returning nil to indicate successful execution of the action
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))       // Sorting the flags of the CLI application by name
	sort.Sort(cli.CommandsByName(app.Commands)) // Sorting the commands of the CLI application by name

	// Running the CLI application with the command-line arguments and handling any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err) // Logging fatal errors
	}
}
