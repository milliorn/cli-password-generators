package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

// Define character sets
const (
	upperChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars  = "abcdefghijklmnopqrstuvwxyz"
	numberChars = "0123456789"
	symbolChars = "!@#$%^&*_-+="
	alphaChars  = upperChars + lowerChars
)

// main is the entry point of the program.
func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	// The app variable represents the command-line application.
	app := &cli.App{
		Name:    "create-password",
		Version: "1.0.0",
		Usage:   "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.",

		// Define the flags that the application accepts.
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "length",
				Aliases: []string{"l"},
				Value:   8,
				Usage:   "Length of the password",
			},

			&cli.BoolFlag{
				Name:    "no-numbers",
				Aliases: []string{"nn"},
				Value:   false,
				Usage:   "Exclude numbers from the password",
			},

			&cli.BoolFlag{
				Name:    "no-symbols",
				Aliases: []string{"ns"},
				Value:   false,
				Usage:   "Exclude special symbols from the password",
			},

			&cli.BoolFlag{
				Name:    "no-uppercase",
				Aliases: []string{"nu"},
				Value:   false,
				Usage:   "Exclude uppercase letters from the password",
			},

			&cli.BoolFlag{
				Name:    "no-lowercase",
				Aliases: []string{"nl"},
				Value:   false,
				Usage:   "Exclude lowercase letters from the password",
			},

			&cli.BoolFlag{
				Name:    "no-letters",
				Aliases: []string{"nle"},
				Value:   false,
				Usage:   "Exclude letters from the password",
			},
		},

		// Define the action that the application should take when it is run.
		Action: func(cCtx *cli.Context) error {
			// Get the values of the flags
			length := cCtx.Int("length")
			noLetters := cCtx.Bool("no-letters")
			noLowercase := cCtx.Bool("no-lowercase")
			noNumbers := cCtx.Bool("no-numbers")
			noSymbols := cCtx.Bool("no-symbols")
			noUppercase := cCtx.Bool("no-uppercase")

			// Get the character set
			chars, err := getCharacterSet(noLetters, noLowercase, noNumbers, noSymbols, noUppercase)

			// If there is an error, return it
			if err != nil {
				return fmt.Errorf("failed to get character set: %w", err)
			}

			// Create a new random number generator
			randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

			// Generate a password
			password := generatePassword(length, randSeed, chars)

			// Print the password
			fmt.Println(password)

			// Return nil to indicate that the action executed successfully
			return nil
		},
	}

	// Sorting flags and commands
	sortFlagsAndCommands(app)

	// Run the application
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error: %s", err)
	}
}

// getCharacterSet returns the character set to use for generating a password.
// If all character types are excluded, an error is returned.
// If no characters are included, an error is returned.
// Otherwise, the character set is returned.
// The character set is a string that contains the characters to use for generating a password.
// The string may contain uppercase and lowercase letters, numbers, and special symbols.
// The noLetters, noLowercase, noNumbers, noSymbols, and noUppercase parameters are used to exclude character types from the character set.
// If a parameter is true, the corresponding character type is excluded from the character set.
// If a parameter is false, the corresponding character type is included in the character set.
// The function returns the character set and an error.
func getCharacterSet(noLetters, noLowercase, noNumbers, noSymbols, noUppercase bool) (string, error) {
	if noLetters && noNumbers && noSymbols {
		return "", errors.New("all character types are excluded, unable to generate password")
	}

	var chars string

	// If no letters are excluded, add them to the character set
	if !noLetters {
		// If no uppercase or lowercase letters are excluded, add them to the character set
		if !noLowercase && !noUppercase {
			chars += alphaChars
		} else {
			// If no letters are excluded, add them to the character set
			if !noLowercase {
				chars += lowerChars
			}
			// If no letters are excluded, add them to the character set
			if !noUppercase {
				chars += upperChars
			}
		}
	}

	// If no numbers or symbols are excluded, add them to the character set
	if !noNumbers {
		chars += numberChars
	}

	// If no symbols are excluded, add them to the character set
	if !noSymbols {
		chars += symbolChars
	}

	// If no characters are included, return an error
	if chars == "" {
		// If no characters are included, return an error
		return "", errors.New("no characters available for password generation")
	}

	// If no characters are included, return an error
	return chars, nil
}

// generatePassword returns a randomly generated password.
// The length parameter specifies the length of the password.
// The randSeed parameter is a random number generator.
// The chars parameter is the character set to use for generating the password.
// The function returns the generated password.
// The password is a string that contains the randomly generated password.
// The password length is equal to the length parameter.
func generatePassword(length int, randSeed *rand.Rand, chars string) string {
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = chars[randSeed.Intn(len(chars))]
	}

	return string(password)
}

// sortFlagsAndCommands sorts the flags and commands of the application.
// The app parameter is the application to sort.
func sortFlagsAndCommands(app *cli.App) {
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}
