# CLI Password Generator

A simple yet powerful command-line interface (CLI) tool for generating secure passwords. This Python script allows users to customize the length and composition of passwords, choosing to include or exclude uppercase letters, lowercase letters, numbers, and symbols. Additionally, users can specify an output file to save the generated password.

## Features

- **Customizable Length:** Set the desired password length.
- **Character Set Selection:** Include or exclude specific character sets (uppercase, lowercase, numbers, symbols).
- **Output Options:** Print the password to the console or save it to a file.
- **Error Handling:** Comprehensive error checks to ensure valid password generation configurations.

## Requirements

- **Python 3.6 or higher**

## Usage

Run the script from the command line, specifying your options. Here are some example commands:

```bash
# Generate a default password (8 characters long, all character types included)
python password_generator.py

# Generate a 12-character long password and exclude symbols
python password_generator.py --length 12 --no-symbols

# Generate a password and save it to a specified file
python password_generator.py -o mypassword.txt

# Display help for more options
python password_generator.py -h
```

You can do the same with the .exe included by just doing `./password_generator.exe`

## Options

- **-l, --length:** Specify the length of the password (default: 8).
- **-nl, --no-lowercase:** Exclude lowercase letters from the password.
- **-nle, --no-letters:** Exclude all letters from the password.
- **-nn, --no-numbers:** Exclude numbers from the password.
- **-ns, --no-symbols:** Exclude special symbols from the password.
- **-nu, --no-uppercase:** Exclude uppercase letters from the password.
- **-o, --output:** Output file to write the password to (if not specified, prints to console).

## Testing

The script comes with a comprehensive test suite to ensure functionality. To run the tests, navigate to the project directory and execute:

```bash
python -m unittest
```

or

```bash
python -m unittest -v
```

This command will run all defined tests, checking the correctness of the password generator under various conditions.

## Contributing

Contributions are welcome! If you have suggestions for improvements or have found a bug, please feel free to create an issue or submit a pull request.

