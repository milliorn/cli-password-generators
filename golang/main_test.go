package main

import (
	"math/rand"
	"strings"
	"testing"
	"time"

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
		{"TestNoNumbersFlagSet", []string{"--no-numbers"}, 0},
		{"TestNoSymbolsFlagSet", []string{"--no-symbols"}, 0},
		{"TestNoUppercaseFlagSet", []string{"--no-uppercase"}, 0},
		{"TestNoLowercaseFlagSet", []string{"--no-lowercase"}, 0},
		{"TestNoLettersFlagSet", []string{"--no-letters"}, 0},
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

func TestGetCharacterSet(t *testing.T) {
	// Test cases covering different scenarios
	tests := []struct {
		name                 string
		noLetters            bool
		noLowercase          bool
		noNumbers            bool
		noSymbols            bool
		noUppercase          bool
		expectedCharacterSet string
		expectedError        error
	}{
		{
			name:                 "AllCharactersIncluded",
			noLetters:            false,
			noLowercase:          false,
			noNumbers:            false,
			noSymbols:            false,
			noUppercase:          false,
			expectedCharacterSet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*_-+=",
			expectedError:        nil,
		},
		{
			name:                 "NoLetters",
			noLetters:            true,
			noLowercase:          false,
			noNumbers:            false,
			noSymbols:            false,
			noUppercase:          false,
			expectedCharacterSet: "0123456789!@#$%^&*_-+=",
			expectedError:        nil,
		},
		{
			name:                 "NoLowercase",
			noLetters:            false,
			noLowercase:          true,
			noNumbers:            false,
			noSymbols:            false,
			noUppercase:          false,
			expectedCharacterSet: "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*_-+=",
			expectedError:        nil,
		},
		{
			name:                 "NoNumbers",
			noLetters:            false,
			noLowercase:          false,
			noNumbers:            true,
			noSymbols:            false,
			noUppercase:          false,
			expectedCharacterSet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*_-+=",
			expectedError:        nil,
		},
		{
			name:                 "NoSymbols",
			noLetters:            false,
			noLowercase:          false,
			noNumbers:            false,
			noSymbols:            true,
			noUppercase:          false,
			expectedCharacterSet: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
			expectedError:        nil,
		},
		{
			name:                 "NoUppercase",
			noLetters:            false,
			noLowercase:          false,
			noNumbers:            false,
			noSymbols:            false,
			noUppercase:          true,
			expectedCharacterSet: "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*_-+=",
			expectedError:        nil,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Call the getCharacterSet function with test input parameters
			characterSet, err := getCharacterSet(test.noLetters, test.noLowercase, test.noNumbers, test.noSymbols, test.noUppercase)

			// Check if the returned character set matches the expected character set
			if characterSet != test.expectedCharacterSet {
				t.Errorf("Expected character set %s, but got %s", test.expectedCharacterSet, characterSet)
			}

			// Check if the returned error matches the expected error
			if (err == nil && test.expectedError != nil) || (err != nil && test.expectedError == nil) || (err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error()) {
				t.Errorf("Expected error %v, but got %v", test.expectedError, err)
			}
		})
	}
}

func TestGeneratePassword(t *testing.T) {
	// Test cases covering different scenarios
	tests := []struct {
		name   string
		length int
		chars  string
	}{
		{
			name:   "Length8",
			length: 8,
			chars:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*_-+=",
		},
		{
			name:   "Length16",
			length: 16,
			chars:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*_-+=",
		},
		{
			name:   "Length32",
			length: 32,
			chars:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*_-+=",
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a new random number generator
			randSeed := rand.New(rand.NewSource(time.Now().UnixNano()))

			// Call the generatePassword function
			password := generatePassword(test.length, randSeed, test.chars)

			// Check if the length of the generated password matches the expected length
			if len(password) != test.length {
				t.Errorf("Expected password length %d, but got %d", test.length, len(password))
			}

			// Check if each character in the generated password belongs to the provided character set
			for _, char := range password {
				if !strings.ContainsRune(test.chars, char) {
					t.Errorf("Generated password contains invalid character: %c", char)
				}
			}
		})
	}
}
