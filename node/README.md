# Node.js / TypeScript Password Generator

A command-line password generator built with Node.js and TypeScript. Generates a random password, copies it to the system clipboard automatically, and optionally saves it to a file.

## Prerequisites

- **Node.js 18 or later**
- **npm** (bundled with Node.js)

## Installation

```bash
cd node
npm install
npm run rebuild
```

`npm run rebuild` erases `dist/` and recompiles the TypeScript source. Re-run it any time you modify files under `src/`.

## Usage

```bash
node ./dist/index [options]
```

Running with no options generates an 8-character password, copies it to the clipboard, and prints a confirmation.

## Flags

| Flag                | Short | Default | Description                            |
| ------------------- | ----- | ------- | -------------------------------------- |
| `--length <number>` | `-l`  | `8`     | Total length of the generated password |
| `--save`            | `-s`  | `false` | Append the password to `passwords.txt` |
| `--no-numbers`      | `-n`  | `false` | Exclude digits                         |
| `--no-symbols`      | `-y`  | `false` | Exclude special symbols                |
| `--version`         | n/a   | n/a     | Print the application version and exit |
| `--help`            | `-h`  | n/a     | Print help text and exit               |

> Every invocation copies the password to the clipboard via `clipboardy`, regardless of other flags.

## Examples

```bash
# Default: 8-character password (copied to clipboard)
node ./dist/index

# 20-character password, saved to passwords.txt
node ./dist/index --length 20 --save

# Letters only (no digits, no symbols), saved to passwords.txt
node ./dist/index --no-numbers --no-symbols --save
```

### npm convenience scripts

```bash
npm run save           # Generate and save to passwords.txt
npm run no-numbers     # Exclude digits, save to passwords.txt
npm run no-symbols     # Exclude symbols, save to passwords.txt
npm run only-letters   # Exclude both digits and symbols, save to passwords.txt
```

## Dependencies

| Package       | Purpose                                     |
| ------------- | ------------------------------------------- |
| `commander`   | CLI argument parsing                        |
| `chalk`       | Terminal output colouring                   |
| `clipboardy`  | Copies the password to the system clipboard |
| `random-seed` | Seedable pseudo-random number generator     |
