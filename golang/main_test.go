package main

import "testing"

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
