# node-cli-password-generator

This repository contains a simple command-line application for generating passwords. It is built using Node.js and TypeScript and utilizes various libraries to generate random passwords and perform other functionalities.

## Installation

To install the dependencies, run the following command `npm install`

## Usage

To run the application and generate a password, use the following command `npm run save`

This will generate a password with a default length of 8 characters and save it to the `passwords.txt` file.

You can also customize the behavior of the password generation using the following options:

- `-l, --length`: Specify the length of the password (default: 8).
- `-s, --save`: Save the generated password to the `passwords.txt` file.
- `-n, --no-numbers`: Remove numbers from the generated password.
- `-y, --no-symbols`: Remove symbols from the generated password.

For example, to generate a password with a length of 20 characters and save it to the `passwords.txt` file, use the following command `node ./dist/index --length=20 -s`

## Examples

Generate a 20-character password and save it:
node ./dist/index --length=20 --save


## Dependencies

The application relies on the following dependencies:

- [chalk](https://www.npmjs.com/package/chalk): Library for styling the command-line output.
- [clipboardy](https://www.npmjs.com/package/clipboardy): Library for accessing the system clipboard.
- [commander](https://www.npmjs.com/package/commander): Library for building command-line interfaces.
- [random-seed](https://www.npmjs.com/package/random-seed): Library for generating random numbers using a seed.
