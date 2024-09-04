import unittest
from password_generator import create_parser, get_character_set, generate_password, NUMBER_CHARS, SYMBOL_CHARS


class TestPasswordGenerator(unittest.TestCase):
    def test_parser_defaults(self):
        parser = create_parser()
        # Simulate command line argument parsing with no arguments
        args = parser.parse_args([])
        self.assertEqual(args.length, 8)
        self.assertFalse(args.no_lowercase)
        self.assertFalse(args.no_letters)
        self.assertFalse(args.no_numbers)
        self.assertFalse(args.no_symbols)
        self.assertFalse(args.no_uppercase)

    def test_character_set_restriction(self):
        # This test checks that excluding letters but not numbers or symbols results in a valid character set of just numbers and symbols.
        chars = get_character_set(no_letters=True, no_lowercase=True,
                                  no_numbers=False, no_symbols=False, no_uppercase=True)
        # Ensure these are defined or imported
        expected_chars = NUMBER_CHARS + SYMBOL_CHARS
        self.assertEqual(set(chars), set(expected_chars),
                         "Character set should only contain numbers and symbols")

    def test_generate_password_length(self):
        chars = "abcdef"
        password = generate_password(4, chars)
        self.assertEqual(len(password), 4)

    def test_password_content(self):
        chars = "abcdef"
        password = generate_password(4, chars)
        for char in password:
            self.assertIn(char, chars)

    def test_password_uniqueness(self):
        chars = "abcdef"
        passwords = set(generate_password(6, chars) for _ in range(100))
        # Assuming a high enough number to statistically ensure uniqueness
        self.assertTrue(len(passwords) > 50)

    def test_all_character_types_excluded(self):
        # This should raise an error because all character types are excluded
        with self.assertRaises(ValueError):
            get_character_set(no_letters=True, no_lowercase=True,
                              no_numbers=True, no_symbols=True, no_uppercase=True)

    def test_no_letters(self):
        # This should not raise an error because numbers and symbols can still be used.
        chars = get_character_set(True, False, False, False, False)
        self.assertTrue(set(chars).issubset(set(NUMBER_CHARS + SYMBOL_CHARS)))

    def test_character_set_restriction_with_value_error(self):
        # Test with all character types excluded
        with self.assertRaises(ValueError):
            get_character_set(no_letters=True, no_lowercase=True,
                              no_numbers=True, no_symbols=True, no_uppercase=True)


if __name__ == '__main__':
    unittest.main()
