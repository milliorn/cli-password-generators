package main

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestCharacterSets(t *testing.T) {
	// Define expected values
	expectedUpperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	expectedLowerChars := "abcdefghijklmnopqrstuvwxyz"
	expectedNumberChars := "0123456789"
	expectedSymbolChars := "!@#$%^&*_-+="
	expectedAlphaChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Test upperChars
	if upperChars != expectedUpperChars {
		t.Errorf("Expected upperChars to be %s, but got %s", expectedUpperChars, upperChars)
	}

	// Test lowerChars
	if lowerChars != expectedLowerChars {
		t.Errorf("Expected lowerChars to be %s, but got %s", expectedLowerChars, lowerChars)
	}

	// Test numberChars
	if numberChars != expectedNumberChars {
		t.Errorf("Expected numberChars to be %s, but got %s", expectedNumberChars, numberChars)
	}

	// Test symbolChars
	if symbolChars != expectedSymbolChars {
		t.Errorf("Expected symbolChars to be %s, but got %s", expectedSymbolChars, symbolChars)
	}

	// Test alphaChars
	if alphaChars != expectedAlphaChars {
		t.Errorf("Expected alphaChars to be %s, but got %s", expectedAlphaChars, alphaChars)
	}
}

func TestFlags(t *testing.T) {
	// Define the app with flags
	app := &cli.App{
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
	}

	// Test cases for each flag
	tests := []struct {
		name     string
		args     []string
		expected int // Define expected values here based on your test cases
	}{
		{"TestLengthFlagSet", []string{"--length=10"}, 10},
		{"TestNoNumbersFlagSet", []string{"--no-numbers"}, 0},     // Assuming it's 0 if set
		{"TestNoSymbolsFlagSet", []string{"--no-symbols"}, 0},     // Assuming it's 0 if set
		{"TestNoUppercaseFlagSet", []string{"--no-uppercase"}, 0}, // Assuming it's 0 if set
		{"TestNoLowercaseFlagSet", []string{"--no-lowercase"}, 0}, // Assuming it's 0 if set
		{"TestNoLettersFlagSet", []string{"--no-letters"}, 0},     // Assuming it's 0 if set
	}
	// Iterate through the test cases
	for _, tt := range tests {
		// Run sub-test
		t.Run(tt.name, func(t *testing.T) {
			// Parse flags
			err := app.Run(tt.args)
			// Check if the error is not nil
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

