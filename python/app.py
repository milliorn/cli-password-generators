import argparse
import os

APP_NAME = "CLI-Password Generator"
LOWER_CHARS = "abcdefghijklmnopqrstuvwxyz"
UPPER_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
NUMBER_CHARS = "0123456789"
SYMBOL_CHARS = "!@#$%^&*_-+="
ALPHA_CHARS = UPPER_CHARS + LOWER_CHARS
DEFAULT_LENGTH = 8
DEFAULT_BOOL = False


def create_parser():
    parser = argparse.ArgumentParser(description=APP_NAME)

    parser.add_argument(
        "-l", "--length", type=int, help="Length of the password", default=DEFAULT_LENGTH)
    parser.add_argument("-nn", "--no-numbers",
                        help="Exclude numbers from the password", action='store_true', default=DEFAULT_BOOL)
    parser.add_argument("-ns", "--no-symbols",
                        help="Exclude special symbols from the password", action='store_true', default=DEFAULT_BOOL)
    parser.add_argument("-nu", "--no-uppercase",
                        help="Exclude uppercase letters from the password", action='store_true', default=DEFAULT_BOOL)
    parser.add_argument("-nl", "--no-lowercase", dest='no_lowercase',
                        help="Exclude lowercase letters from the password", action='store_true', default=DEFAULT_BOOL)
    parser.add_argument("-nle", "--no-letters", dest='no_letters',
                        help="Exclude letters from the password", action='store_true', default=DEFAULT_BOOL)

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


def main():
    parser = create_parser()
    args = parser.parse_args()

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

    # print(chars)


if __name__ == "__main__":
    main()
