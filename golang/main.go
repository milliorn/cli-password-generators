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

const alphaLower = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const symbols = "!@#$%^&*_-+="

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

		// Defining the flags of the CLI application
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "override-length",                     // Setting the name of the flag
				Aliases: []string{"ol"},                        // Setting the aliases of the flag
				Value:   false,                                 // Setting the default value of the flag
				Usage:   "Override the length of the password", // Setting the usage description of the flag
			},

			&cli.IntFlag{
				Name:    "length",                 // Setting the name of the flag
				Aliases: []string{"l"},            // Setting the aliases of the flag
				Value:   8,                        // Setting the default value of the flag
				Usage:   "Length of the password", // Setting the usage description of the flag
				Action: func(ctx *cli.Context, v int) error { // Defining the action to be executed when the flag is set
					if v > 74 && !ctx.Bool("override-length") {
						return fmt.Errorf("length of password must be less than 74 due to sample size of characters available. Please try again with a smaller length")
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
				Name:    "no-letters",                        // Setting the name of the flag
				Aliases: []string{"nle"},                     // Setting the aliases of the flag
				Value:   false,                               // Setting the default value of the flag
				Usage:   "Exclude letters from the password", // Setting the usage description of the flag
			},
		},

		Action: func(cCtx *cli.Context) error { // Defining the action to be executed when the CLI application is run
			// Getting the value of flags from the context
			length := cCtx.Int("length")
			noLetters := cCtx.Bool("no-letters")
			noLowercase := cCtx.Bool("no-lowercase")
			noNumbers := cCtx.Bool("no-numbers")
			noSymbols := cCtx.Bool("no-symbols")
			noUppercase := cCtx.Bool("no-uppercase")

			chars := "" // Defining a variable to store the characters to be used in the password

			// Checking if letters are not excluded
			// Checking if lowercase letters are not excluded
			// Checking if uppercase letters are not excluded
			chars = generateAlpha(noLetters, noLowercase, chars, noUppercase)

			if !noNumbers { // Checking if numbers are not excluded
				chars += numbers
			}

			if !noSymbols { // Checking if symbols are not excluded
				chars += symbols
			}

			randSeed := rand.New(rand.NewSource(time.Now().UnixNano())) // Creating a new random seed

			var password string // Defining a variable to store the password

			// Generating the password
			for i := 0; i < length; i++ {
				n := randSeed.Intn(len(chars)) // Generating a random number within the length of the characters
				password += string(chars[n])   // Adding the character at the random index to the password
			}

			if length > len(password) { // Checking if the length of the password requested is greater than the length of the password generated
				log.Fatal("length of password requested is greater than the length of the password generated. Please try again.")
			}

			fmt.Println(password)
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))       // Sorting the flags of the CLI application by name
	sort.Sort(cli.CommandsByName(app.Commands)) // Sorting the commands of the CLI application by name

	// Running the CLI application with the command-line arguments and handling any errors
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err) // Logging fatal errors
	}
}

// generateAlpha generates an alphanumeric string based on the provided parameters.
// If noLetters is false, it includes both lowercase and uppercase alphabets in the generated string.
// If noLowercase is false, it includes lowercase alphabets in the generated string.
// If noUppercase is false, it includes uppercase alphabets in the generated string.
// The chars parameter is used to append additional characters to the generated string.
// The function returns the generated alphanumeric string.
func generateAlpha(noLetters bool, noLowercase bool, chars string, noUppercase bool) string {
	if !noLetters {
		if !noLowercase {
			chars += alphaLower
		}
		if !noUppercase {
			chars += alphaUpper
		}
	}
	return chars
}
