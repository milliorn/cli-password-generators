package main

import (
	"fmt" // Importing the 'fmt' package for formatted I/O operations
	"log" // Importing the 'log' package for logging errors
	"os"  // Importing the 'os' package for accessing command-line arguments and environment variables
	"sort"

	"github.com/urfave/cli/v2" // Importing the 'github.com/urfave/cli/v2' package for building CLI applications
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{ // Setting the version flag of the CLI application
		Name:    "version",                // Setting the name of the flag
		Aliases: []string{"v"},            // Setting the aliases of the flag
		Usage:   "print only the version", // Setting the usage description of the flag
	}

	// Creating a new CLI application instance with name and usage description
	app := &cli.App{
		Name:    "create-password",                                                                                                      // Setting the name of the CLI application
		Version: "1.0.0",                                                                                                                // Setting the version of the CLI application
		Usage:   "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.", // Setting the usage description of the CLI application
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "length",                 // Setting the name of the flag
				Value: 8,                        // Setting the default value of the flag
				Usage: "Length of the password", // Setting the usage description of the flag
				Action: func(ctx *cli.Context, v int) error { // Defining the action to be executed when the flag is set
					if v > 95 {
						return fmt.Errorf("length of password must be less than 95")
					}
					return nil // Returning nil to indicate successful execution of the action
				},
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
		Action: func(cCtx *cli.Context) error { // Defining the action to be executed when the CLI application is run
			// fmt.Println("Hello, World!") // Printing "Hello, World!" to the console
			createPassword(cCtx) // Calling the 'createPassword' function with the context of the CLI application
			return nil           // Returning nil to indicate successful execution of the action
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))       // Sorting the flags of the CLI application by name
	sort.Sort(cli.CommandsByName(app.Commands)) // Sorting the commands of the CLI application by name

	// Running the CLI application with the command-line arguments and handling any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err) // Logging fatal errors
	}
}

func createPassword(cCtx *cli.Context) {
	length := cCtx.Int("length")                // Getting the value of the 'length' flag
	no_numbers := cCtx.Bool("no-numbers")       // Getting the value of the 'no-numbers' flag
	no_symbols := cCtx.Bool("no-symbols")       // Getting the value of the 'no-symbols' flag
	no_uppercase := cCtx.Bool("no-uppercase")   // Getting the value of the 'no-uppercase' flag
	no_lowercase := cCtx.Bool("no-lowercase")   // Getting the value of the 'no-lowercase' flag
	no_similar := cCtx.Bool("no-similar")       // Getting the value of the 'no-similar' flag
	no_vowels := cCtx.Bool("no-vowels")         // Getting the value of the 'no-vowels' flag
	no_consonants := cCtx.Bool("no-consonants") // Getting the value of the 'no-consonants' flag
	no_letters := cCtx.Bool("no-letters")       // Getting the value of the 'no-letters' flag

	fmt.Println(length) // Printing the value of the 'length' flag
	fmt.Println(no_numbers)
	fmt.Println(no_symbols)
	fmt.Println(no_uppercase)
	fmt.Println(no_lowercase)
	fmt.Println(no_similar)
	fmt.Println(no_vowels)
	fmt.Println(no_consonants)
	fmt.Println(no_letters)

	panic("unimplemented")
}
