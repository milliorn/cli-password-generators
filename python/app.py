import argparse
import os

from random import shuffle

APP_NAME = "CLI-Password Generator"
DEFAULT_BOOL = False
DEFAULT_LENGTH = 8

NUMBER_CHARS = "0123456789"
SYMBOL_CHARS = "!@#$%^&*_-+="

UPPER_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
LOWER_CHARS = "abcdefghijklmnopqrstuvwxyz"

ALPHA_CHARS = UPPER_CHARS + LOWER_CHARS


def create_parser():
    parser = argparse.ArgumentParser(description=APP_NAME)

    parser.add_argument("-l", "--length", type=int,
                        help="Length of the password", default=DEFAULT_LENGTH)
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
    if no_letters and no_numbers and no_symbols:
        raise ValueError(
            "All character types are excluded, unable to generate password")

    buffer_character_set = []

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

    if len(buffer_character_set) == 0:
        raise ValueError("No characters available for password generation")

    character_set = "".join(buffer_character_set)

    return character_set


def generate_password(length, chars):
    # Convert the string of characters into a list
    char_list = list(chars)
    # Shuffle the list of characters
    shuffle(char_list)
    # Join the list back into a string and select the first 'length' characters for the password
    password = ''.join(char_list)[:length]

    return password


def main():
    parser = create_parser()
    args = parser.parse_args()

    if args.length < 1:
        raise ValueError(
            f"Length must be greater than 0. Provided length: {args.length}")

    try:
        # Get the character set
        chars = get_character_set(
            args.no_letters,
            args.no_lowercase,
            args.no_numbers,
            args.no_symbols,
            args.no_uppercase)

    except ValueError as e:
        # Handle the error by re-raising it with additional context
        raise ValueError(f"Failed to get character set: {e}")

    password = generate_password(args.length, chars)
    print(password)


if __name__ == "__main__":
    main()
