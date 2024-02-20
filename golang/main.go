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
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	numberChars  = "0123456789"
	symbolChars  = "!@#$%^&*_-+="
	defaultChars = upperChars + lowerChars + numberChars + symbolChars
)

var alphaChars = upperChars + lowerChars

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "create-password",
		Version: "1.0.0",
		Usage:   "Generate a password involves creating a random mix of uppercase and lowercase letters, numbers, and special symbols.",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "override-length",
				Aliases: []string{"ol"},
				Value:   false,
				Usage:   "Override the length of the password",
			},

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

		Action: func(cCtx *cli.Context) error {
			length := cCtx.Int("length")
			noLetters := cCtx.Bool("no-letters")
			noLowercase := cCtx.Bool("no-lowercase")
			noNumbers := cCtx.Bool("no-numbers")
			noSymbols := cCtx.Bool("no-symbols")
			noUppercase := cCtx.Bool("no-uppercase")

			chars, err := getCharacterSet(noLetters, noLowercase, noNumbers, noSymbols, noUppercase)

			if err != nil {
				return fmt.Errorf("failed to get character set: %w", err)
			}

			randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

			password := generatePassword(length, randSeed, chars)

			fmt.Println(password)
			return nil
		},
	}

	// Sorting flags and commands
	sortFlagsAndCommands(app)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func getCharacterSet(noLetters, noLowercase, noNumbers, noSymbols, noUppercase bool) (string, error) {
	if noLetters && noNumbers && noSymbols {
		return "", errors.New("all character types are excluded, unable to generate password")
	}

	var chars string

	if !noLetters {
		if !noLowercase && !noUppercase {
			chars += alphaChars
		} else {
			if !noLowercase {
				chars += lowerChars
			}
			if !noUppercase {
				chars += upperChars
			}
		}
	}
	if !noNumbers {
		chars += numberChars
	}
	if !noSymbols {
		chars += symbolChars
	}

	if chars == "" {
		// If no characters are included, return an error
		return "", errors.New("no characters available for password generation")
	}

	return chars, nil
}

func generatePassword(length int, randSeed *rand.Rand, chars string) string {
	password := make([]byte, length)

	for i := 0; i < length; i++ {
		password[i] = chars[randSeed.Intn(len(chars))]
	}

	return string(password)
}

func sortFlagsAndCommands(app *cli.App) {
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
}
