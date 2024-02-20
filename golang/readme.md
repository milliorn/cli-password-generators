# Go Password Generator

This is a simple command-line password generator written in Go. It generates passwords by creating a random mix of uppercase and lowercase letters, numbers, and special symbols.

## Usage

To use the password generator, follow these steps:

1. Build the program using the `go build` command.
2. Run the program using the `./password-generator` command.

## Example

Here is an example of using the password generator:

```bash
./password-generator --l 12 --nn
```

This command generates a 12-character password without numbers.

## Flags

The password generator supports the following flags:

- `--length`, `-l`: Length of the password (default is 8).
- `--no-numbers`, `-nn`: Exclude numbers from the password.
- `--no-symbols`, `-ns`: Exclude special symbols from the password.
- `--no-uppercase`, `-nu`: Exclude uppercase letters from the password.
- `--no-lowercase`, `-nl`: Exclude lowercase letters from the password.
- `--no-letters`, `-nle`: Exclude letters from the password.

## Tests

To run the tests, use the `go test` command.
