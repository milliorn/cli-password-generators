import argparse
import os
from random import shuffle

# Constants for the application
APP_NAME = "CLI-Password Generator"
DEFAULT_BOOL = False
DEFAULT_LENGTH = 8

# Character sets used in password generation
NUMBER_CHARS = "0123456789"
SYMBOL_CHARS = "!@#$%^&*_-+="
UPPER_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
LOWER_CHARS = "abcdefghijklmnopqrstuvwxyz"
ALPHA_CHARS = UPPER_CHARS + LOWER_CHARS


def create_parser():
    # Creates and configures the argument parser with options for the password generator.
    parser = argparse.ArgumentParser(description=APP_NAME)

    # Add command line options for configuring the password
    parser.add_argument("-l", "--length", type=int, default=DEFAULT_LENGTH,
                        help="Specify the length of the password (default: 8).")

    parser.add_argument("-nl", "--no-lowercase", dest='no_lowercase',
                        help="Exclude lowercase letters from the password", action='store_true', default=DEFAULT_BOOL)

    parser.add_argument("-nle", "--no-letters", dest='no_letters',
                        help="Exclude letters from the password", action='store_true', default=DEFAULT_BOOL)

    parser.add_argument("-nn", "--no-numbers", help="Exclude numbers from the password",
                        action='store_true', default=DEFAULT_BOOL)

    parser.add_argument("-ns", "--no-symbols", help="Exclude special symbols from the password",
                        action='store_true', default=DEFAULT_BOOL)

    parser.add_argument("-nu", "--no-uppercase", help="Exclude uppercase letters from the password",
                        action='store_true', default=DEFAULT_BOOL)

    return parser


def get_character_set(no_letters, no_lowercase, no_numbers, no_symbols, no_uppercase):
    # Determines the set of characters that can be used in the password based on the flags provided.
    if no_letters and no_numbers and no_symbols:
        # If all character types are excluded, raise an error
        raise ValueError(
            "All character types are excluded, unable to generate password")

    buffer_character_set = []

    # Determine which character sets to include based on the flags
    if not no_letters:
        if not no_lowercase and not no_uppercase:
            buffer_character_set.append(ALPHA_CHARS)
        else:
            if not no_lowercase:
                buffer_character_set.append(LOWER_CHARS)
            if not no_uppercase:
                buffer_character_set.append(UPPER_CHARS)

    if not no_numbers:
        buffer_character_set.append(NUMBER_CHARS)

    if not no_symbols:
        buffer_character_set.append(SYMBOL_CHARS)

    # If no characters are available after checking all conditions, raise an error
    if len(buffer_character_set) == 0:
        raise ValueError("No characters available for password generation")

    # Combine all the characters into one string
    character_set = "".join(buffer_character_set)

    return character_set


def generate_password(length, chars):
    # Generates a password of specified length from the given set of characters.
    if length < 1:
        # Ensure the password length is at least 1
        raise ValueError(f"Length must be at least 1. Got {length}")

    # Shuffle the characters to ensure randomness
    char_list = list(chars)
    shuffle(char_list)

    # Generate the password by selecting the first 'length' characters
    password = ''.join(char_list)[:length]

    return password


def main():
    parser = create_parser()
    args = parser.parse_args()

    # Check if the provided length is valid
    if args.length < 1:
        raise ValueError(
            f"Length must be greater than 0. Provided length: {args.length}")

    try:
        # Obtain the set of characters to use for the password based on user input
        chars = get_character_set(
            args.no_letters,
            args.no_lowercase,
            args.no_numbers,
            args.no_symbols,
            args.no_uppercase)

    except ValueError as e:
        # Handle errors related to character set generation
        raise ValueError(f"Failed to get character set: {e}")

    # Generate and print the password
    password = generate_password(args.length, chars)
    print(password)


if __name__ == "__main__":
    # Entry point of the program when run as a script
    main()
