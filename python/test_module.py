import unittest
from password_generator import create_parser, get_character_set, generate_password, NUMBER_CHARS, SYMBOL_CHARS, ALPHA_CHARS, LOWER_CHARS, UPPER_CHARS


class TestPasswordGenerator(unittest.TestCase):
    def test_parser_defaults(self):
        # Test if the command line argument parser sets the correct default values
        parser = create_parser()
        # Simulate parsing with no arguments provided
        args = parser.parse_args([])
        self.assertEqual(args.length, 8)  # Default password length should be 8
        # Check that all character type exclusions are False by default
        self.assertFalse(args.no_lowercase)
        self.assertFalse(args.no_letters)
        self.assertFalse(args.no_numbers)
        self.assertFalse(args.no_symbols)
        self.assertFalse(args.no_uppercase)

    def test_character_set_restriction(self):
        # Test that excluding letters (both upper and lower) but allowing numbers and symbols
        # results in a character set that contains only numbers and symbols.
        chars = get_character_set(no_letters=True, no_lowercase=True,
                                  no_numbers=False, no_symbols=False, no_uppercase=True)
        expected_chars = NUMBER_CHARS + SYMBOL_CHARS  # Expected character set
        self.assertEqual(set(chars), set(expected_chars),
                         "Character set should only contain numbers and symbols")

    def test_generate_password_length(self):
        # Verify that the password generated is of the specified length
        chars = "abcdef"
        password = generate_password(4, chars)
        self.assertEqual(len(password), 4)

    def test_password_content(self):
        # Check that the generated password contains only characters from the specified set
        chars = "abcdef"
        password = generate_password(4, chars)
        for char in password:
            self.assertIn(char, chars)

    def test_password_uniqueness(self):
        # Verify that generating multiple passwords results in a high degree of uniqueness
        chars = "abcdef"
        passwords = set(generate_password(6, chars) for _ in range(100))
        # Expect a significant number of unique passwords
        self.assertTrue(len(passwords) > 50)

    def test_all_character_types_excluded(self):
        # Test that excluding all character types raises a ValueError
        with self.assertRaises(ValueError):
            get_character_set(no_letters=True, no_lowercase=True,
                              no_numbers=True, no_symbols=True, no_uppercase=True)

    def test_no_letters(self):
        # Verify that excluding all letters but allowing numbers and symbols does not raise an error
        chars = get_character_set(True, False, False, False, False)
        self.assertTrue(set(chars).issubset(set(NUMBER_CHARS + SYMBOL_CHARS)))

    def test_character_set_restriction_with_value_error(self):
        # Test again that excluding all character types raises a ValueError
        with self.assertRaises(ValueError):
            get_character_set(no_letters=True, no_lowercase=True,
                              no_numbers=True, no_symbols=True, no_uppercase=True)


class TestFlagCombinations(unittest.TestCase):
    def test_no_lowercase_only(self):
        # Verify that setting the no_lowercase flag excludes lowercase characters
        chars = get_character_set(no_letters=False, no_lowercase=True,
                                  no_numbers=False, no_symbols=False, no_uppercase=False)
        # Lowercase characters should not be present
        self.assertNotIn(LOWER_CHARS, chars)
        # Uppercase characters should be present
        self.assertIn(UPPER_CHARS, chars)

    def test_no_uppercase_only(self):
        # Verify that setting the no_uppercase flag excludes uppercase characters
        chars = get_character_set(no_letters=False, no_lowercase=False,
                                  no_numbers=False, no_symbols=False, no_uppercase=True)
        # Uppercase characters should not be present
        self.assertNotIn(UPPER_CHARS, chars)
        # Lowercase characters should be present
        self.assertIn(LOWER_CHARS, chars)

    def test_no_numbers_no_symbols(self):
        # Verify that excluding numbers and symbols leaves only alphabetic characters
        chars = get_character_set(no_letters=False, no_lowercase=False,
                                  no_numbers=True, no_symbols=True, no_uppercase=False)
        # Numbers and symbols should not be present
        self.assertNotIn(NUMBER_CHARS + SYMBOL_CHARS, chars)
        # Alphabetic characters should be present
        self.assertIn(ALPHA_CHARS, chars)

    def test_no_uppercase_no_lowercase(self):
        # Verify that excluding both uppercase and lowercase characters leaves only numbers and symbols
        chars = get_character_set(no_letters=False, no_lowercase=True,
                                  no_numbers=False, no_symbols=False, no_uppercase=True)
        # No letters should be present
        self.assertNotIn(UPPER_CHARS + LOWER_CHARS, chars)
        # Only numbers and symbols should be present
        self.assertIn(NUMBER_CHARS + SYMBOL_CHARS, chars)

    def test_invalid_length(self):
        # Verify that attempting to generate a password with zero length raises a ValueError
        with self.assertRaises(ValueError):
            generate_password(0, ALPHA_CHARS)

    def test_negative_length(self):
        # Verify that a negative length value raises a ValueError
        with self.assertRaises(ValueError):
            generate_password(-1, ALPHA_CHARS)


if __name__ == '__main__':
    unittest.main()
