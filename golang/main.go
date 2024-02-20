package main

import (
	"fmt" // Importing the 'fmt' package for formatted I/O operations
	"log" // Importing the 'log' package for logging errors
	"math/rand"
	"os" // Importing the 'os' package for accessing command-line arguments and environment variables
	"sort"
	"time"

	"github.com/urfave/cli/v2" // Importing the 'github.com/urfave/cli/v2' package for building CLI applications
)

const alpha_lower = "abcdefghijklmnopqrstuvwxyz"
const alpha_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "!@#$%^&*_-+="

// const vowels = "aeiouAEIOU"
// const consonants = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

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
				Name:    "length",                 // Setting the name of the flag
				Aliases: []string{"l"},            // Setting the aliases of the flag
				Value:   8,                        // Setting the default value of the flag
				Usage:   "Length of the password", // Setting the usage description of the flag
				Action: func(ctx *cli.Context, v int) error { // Defining the action to be executed when the flag is set
					if v > 95 {
						return fmt.Errorf("length of password must be less than 74 due to sample size of characters available. Please try again with a smaller length.")
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

			// &cli.BoolFlag{
			// 	Name:    "no-similar",                                   // Setting the name of the flag
			// 	Aliases: []string{"nsi"},                                // Setting the aliases of the flag
			// 	Value:   false,                                          // Setting the default value of the flag
			// 	Usage:   "Exclude similar characters from the password", // Setting the usage description of the flag
			// },

			// &cli.BoolFlag{
			// 	Name:    "no-vowels",                        // Setting the name of the flag
			// 	Aliases: []string{"nv"},                     // Setting the aliases of the flag
			// 	Value:   false,                              // Setting the default value of the flag
			// 	Usage:   "Exclude vowels from the password", // Setting the usage description of the flag
			// },

			// &cli.BoolFlag{
			// 	Name:    "no-consonants",                        // Setting the name of the flag
			// 	Aliases: []string{"nc"},                         // Setting the aliases of the flag
			// 	Value:   false,                                  // Setting the default value of the flag
			// 	Usage:   "Exclude consonants from the password", // Setting the usage description of the flag
			// },

			&cli.BoolFlag{
				Name:    "no-letters",                        // Setting the name of the flag
				Aliases: []string{"nle"},                     // Setting the aliases of the flag
				Value:   false,                               // Setting the default value of the flag
				Usage:   "Exclude letters from the password", // Setting the usage description of the flag
			},
		},

		Action: func(cCtx *cli.Context) error { // Defining the action to be executed when the CLI application is run
			length := cCtx.Int("length") // Getting the value of the 'length' flag
			// no_consonants := cCtx.Bool("no-consonants") // Getting the value of the 'no-consonants' flag
			no_letters := cCtx.Bool("no-letters")     // Getting the value of the 'no-letters' flag
			no_lowercase := cCtx.Bool("no-lowercase") // Getting the value of the 'no-lowercase' flag
			no_numbers := cCtx.Bool("no-numbers")     // Getting the value of the 'no-numbers' flag
			// no_similar := cCtx.Bool("no-similar")       // Getting the value of the 'no-similar' flag
			no_symbols := cCtx.Bool("no-symbols")     // Getting the value of the 'no-symbols' flag
			no_uppercase := cCtx.Bool("no-uppercase") // Getting the value of the 'no-uppercase' flag
			// no_vowels := cCtx.Bool("no-vowels")       // Getting the value of the 'no-vowels' flag

			chars := "" // Setting the characters to be used for generating the password

			if !no_letters { // Checking if the 'no-letters' flag is set
				if !no_uppercase { // Checking if the 'no-uppercase' flag is not set
					chars += alpha_upper
				}

				if !no_lowercase { // Checking if the 'no-lowercase' flag is not set
					chars += alpha_lower
				}
			}

			if !no_numbers { // Checking if the 'no-numbers' flag is not set
				chars += numbers
			}

			if !no_symbols { // Checking if the 'no-symbols' flag is not set
				chars += symbols
			}

			var password string // Declaring a variable to store the generated password

			randSeed := rand.New(rand.NewSource(time.Now().UnixNano())) // Creating a new random seed

			for i := 0; i < length; i++ { // Generating the password
				n := randSeed.Intn(len(chars)) // Generating a random number within the length of the characters
				password += string(chars[n])   // Appending the character at the random index to the password
			}

			if length > len(password) {
				log.Fatal("length of password requested is greater than the length of the password generated. Please try again with a smaller length.")
			}

			fmt.Println(password) // Printing the generated password
			return nil            // Returning nil to indicate successful execution of the action
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))       // Sorting the flags of the CLI application by name
	sort.Sort(cli.CommandsByName(app.Commands)) // Sorting the commands of the CLI application by name

	// Running the CLI application with the command-line arguments and handling any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err) // Logging fatal errors
	}
}
