# cli-password-generators

[![CodeQL](https://github.com/milliorn/cli-password-generators/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/milliorn/cli-password-generators/actions/workflows/github-code-scanning/codeql)
[![Dependabot reviewer](https://github.com/milliorn/cli-password-generators/actions/workflows/automerge.yml/badge.svg)](https://github.com/milliorn/cli-password-generators/actions/workflows/automerge.yml)
[![Dependabot Updates](https://github.com/milliorn/cli-password-generators/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/milliorn/cli-password-generators/actions/workflows/dependabot/dependabot-updates)

## Overview

This repository contains three functionally equivalent command-line password generator applications, each implemented in a different programming language: **Go**, **Python**, and **Node.js (TypeScript)**. The goal of this project is to demonstrate how the same problem: generating a cryptographically reasonable, configurable random password, can be solved idiomatically across different language ecosystems.

All three implementations share a common design philosophy:

- Accept flags at the command line to control which character classes appear in the generated password.
- Default to a sensible, secure-looking output (mixed case letters, digits, and symbols, 8 characters) so that running the tool with no arguments "just works".
- Fail loudly and helpfully if the user's flag combination would make password generation impossible (e.g., excluding every character class).

Each implementation lives in its own subdirectory and is entirely self-contained. You do not need to install the toolchain for all three languages. Pick the one you are most comfortable with, or use all three if you're comparing behaviour.

---

## Table of Contents

1. [Character Sets](#character-sets)
2. [Go Implementation](#go-implementation)
   - [Prerequisites](#go-prerequisites)
   - [Installation](#go-installation)
   - [Usage](#go-usage)
   - [Flags](#go-flags)
   - [Examples](#go-examples)
   - [Running Tests](#go-tests)
3. [Python Implementation](#python-implementation)
   - [Prerequisites](#python-prerequisites)
   - [Installation](#python-installation)
   - [Usage](#python-usage)
   - [Flags](#python-flags)
   - [Examples](#python-examples)
   - [Running Tests](#python-tests)
4. [Node.js / TypeScript Implementation](#nodejs--typescript-implementation)
   - [Prerequisites](#nodejs-prerequisites)
   - [Installation](#nodejs-installation)
   - [Usage](#nodejs-usage)
   - [Flags](#nodejs-flags)
   - [Examples](#nodejs-examples)
5. [Project Structure](#project-structure)
6. [Contributing](#contributing)
7. [License](#license)

---

## Character Sets

All implementations draw from the same four character pools. Understanding these pools makes it easier to reason about the flags described in the sections below.

| Pool | Characters |
| --- | --- |
| Uppercase letters | `ABCDEFGHIJKLMNOPQRSTUVWXYZ` |
| Lowercase letters | `abcdefghijklmnopqrstuvwxyz` |
| Digits | `0123456789` |
| Symbols | `!@#$%^&*_-+=` |

By default, every pool is included. Flags let you opt individual pools out. The only hard constraint is that at least one pool must remain enabled; attempting to exclude all pools is treated as a usage error.

---

## Go Implementation

### Go Prerequisites

- [Go](https://go.dev/dl/) **1.22 or later**. Verify your installation with:

  ```bash
  go version
  ```

The only external dependency is [`github.com/urfave/cli/v2`](https://github.com/urfave/cli), which is managed automatically by Go modules. You do not need to install it manually.

### Go Installation

```bash
# Clone the repository (skip this step if you already have it)
git clone https://github.com/milliorn/cli-password-generators.git
cd cli-password-generators/golang

# Download module dependencies
go mod download

# Build a native binary
go build -o create-password .
```

After the build succeeds, a binary named `create-password` will be present in the `golang/` directory. You can move it anywhere on your `$PATH` if you want system-wide access.

Alternatively, you can skip the build step entirely and run the program directly through the Go toolchain:

```bash
go run main.go [flags]
```

### Go Usage

```text
create-password [global options]
```

Running the binary with no arguments generates an 8-character password using all four character pools and prints it to standard output.

### Go Flags

| Flag | Short | Type | Default | Description |
| --- | --- | --- | --- | --- |
| `--length` | `-l` | integer | `8` | Total length of the generated password. |
| `--no-letters` | `-nle` | boolean | `false` | Exclude **all** letters (both uppercase and lowercase). |
| `--no-lowercase` | `-nl` | boolean | `false` | Exclude lowercase letters only. |
| `--no-numbers` | `-nn` | boolean | `false` | Exclude digits. |
| `--no-symbols` | `-ns` | boolean | `false` | Exclude special symbols. |
| `--no-uppercase` | `-nu` | boolean | `false` | Exclude uppercase letters only. |
| `--version` | `-v` | boolean | n/a | Print the application version and exit. |
| `--help` | `-h` | boolean | n/a | Print help text and exit. |

> **Note on flag precedence:** `--no-letters` takes priority over `--no-lowercase` and `--no-uppercase`. If you pass `--no-letters`, both letter pools are excluded regardless of the other letter-related flags.

### Go Examples

```bash
# Default: 8-character password with all character classes
./create-password
# Example output: aB3!kZ9@

# 20-character password
./create-password --length 20

# Numbers and symbols only (no letters at all)
./create-password --no-letters

# Uppercase letters and digits only
./create-password --no-lowercase --no-symbols

# 32-character password with only digits and symbols
./create-password --no-letters --length 32

# Print version
./create-password --version
```

### Go Tests

The Go implementation ships with a comprehensive test suite covering character set constants, flag parsing, character set composition logic, and password generation output.

```bash
cd golang
go test ./...
```

To see verbose test output with individual test case names:

```bash
go test -v ./...
```

---

## Python Implementation

### Python Prerequisites

- **Python 3.8 or later**. Verify with:

  ```bash
  python3 --version
  ```

The Python implementation uses **only the standard library** (`argparse`, `random`, `os`, `unittest`). No third-party packages are required, and no virtual environment is necessary.

### Python Installation

```bash
cd cli-password-generators/python
```

That is all. There are no dependencies to install.

### Python Usage

```text
python3 password_generator.py [options]
```

Running the script with no arguments generates an 8-character password and prints it to standard output.

### Python Flags

| Flag | Short | Type | Default | Description |
| --- | --- | --- | --- | --- |
| `--length` | `-l` | integer | `8` | Total length of the generated password. |
| `--no-letters` | `-nle` | boolean | `false` | Exclude **all** letters (both uppercase and lowercase). |
| `--no-lowercase` | `-nl` | boolean | `false` | Exclude lowercase letters only. |
| `--no-numbers` | `-nn` | boolean | `false` | Exclude digits. |
| `--no-symbols` | `-ns` | boolean | `false` | Exclude special symbols. |
| `--no-uppercase` | `-nu` | boolean | `false` | Exclude uppercase letters only. |
| `--output` | `-o` | string/flag | `None` | Write the generated password to a file instead of stdout. When `-o` is passed without a file path argument, the password is written to `password.txt` in the current directory. When a path is supplied (e.g., `-o /tmp/secret.txt`), that file is used. |
| `--help` | `-h` | boolean | n/a | Print help text and exit. |

### Python Examples

```bash
# Default: 8-character password with all character classes
python3 password_generator.py

# 16-character password
python3 password_generator.py --length 16

# Numbers and symbols only
python3 password_generator.py --no-letters

# Uppercase letters and digits only
python3 password_generator.py --no-lowercase --no-symbols

# Save to the default output file (password.txt)
python3 password_generator.py -o

# Save to a specific file
python3 password_generator.py --length 24 --output /tmp/my_password.txt

# 32-character lowercase-only password
python3 password_generator.py --no-uppercase --no-numbers --no-symbols --length 32
```

### Python Tests

The Python test suite uses the built-in `unittest` framework and validates parser defaults, character set construction (including all error conditions), password length, password character validity, and flag combinations.

```bash
cd python
python3 -m unittest test_module.py -v
```

---

## Node.js / TypeScript Implementation

### Node.js Prerequisites

- **Node.js 18 or later**. Verify with:

  ```bash
  node --version
  ```

- **npm** (bundled with Node.js). Verify with:

  ```bash
  npm --version
  ```

The TypeScript source is compiled to plain JavaScript before execution. The compiled output is committed to the `dist/` directory and is ready to run after `npm install`, so you do not need a global TypeScript installation unless you intend to modify the source.

### Node.js Installation

```bash
cd cli-password-generators/node

# Install all runtime and development dependencies
npm install
```

**Key dependencies:**

| Package | Purpose |
| --- | --- |
| `commander` | CLI argument parsing |
| `chalk` | Terminal output colouring |
| `clipboardy` | Automatically copies the generated password to the system clipboard |
| `random-seed` | Seedable pseudo-random number generator |

**Key dev dependencies:** `typescript`, `ts-node`, `nodemon`, `@types/node`, `@types/random-seed`.

If you modify the TypeScript source files under `src/`, recompile with:

```bash
npm run rebuild
```

This erases the existing `dist/` directory and runs `tsc` to produce a fresh build.

### Node.js Usage

```bash
node ./dist/index [options]
```

A set of convenience `npm` scripts are provided for the most common use cases (see below). Running with no options generates an 8-character password, copies it to the clipboard, and prints a confirmation message.

### Node.js Flags

| Flag | Short | Type | Default | Description |
| --- | --- | --- | --- | --- |
| `--length <number>` | `-l` | integer | `8` | Total length of the generated password. |
| `--save` | `-s` | boolean | `false` | Append the generated password to `passwords.txt` in the project directory. |
| `--no-numbers` | `-n` | boolean | `false` | Exclude digits from the password. |
| `--no-symbols` | `-y` | boolean | `false` | Exclude special symbols from the password. |
| `--version` | n/a | boolean | n/a | Print the application version and exit. |
| `--help` | `-h` | boolean | n/a | Print help text and exit. |

> **Note:** The Node.js implementation does not currently expose `--no-uppercase`, `--no-lowercase`, or `--no-letters` flags. The character pool always contains both cases of the alphabet; only digits and symbols can be individually excluded.
>
> **Clipboard behaviour:** Every invocation automatically copies the generated password to the system clipboard via `clipboardy`. The terminal will display `Password copied to clipboard` in yellow text regardless of whether `--save` is used.

### Node.js Examples

```bash
# Default: 8-character password (copied to clipboard automatically)
node ./dist/index

# 20-character password
node ./dist/index --length 20

# Letters only (no digits, no symbols), saved to passwords.txt
node ./dist/index --no-numbers --no-symbols --save

# Using npm convenience scripts
npm run save           # Generate and save to passwords.txt
npm run no-numbers     # Exclude digits, save to passwords.txt
npm run no-symbols     # Exclude symbols, save to passwords.txt
npm run only-letters   # Exclude both digits and symbols, save to passwords.txt
```

---

## Project Structure

```text
cli-password-generators/
├── LICENSE
├── readme.md                  ← You are here
│
├── golang/
│   ├── go.mod                 ← Go module definition (requires Go 1.22+)
│   ├── main.go                ← Application entry point and all logic
│   ├── main_test.go           ← Unit tests
│   ├── readme.md              ← Go-specific notes
│   └── examples/              ← Shell scripts demonstrating various flag combinations
│       ├── run.sh
│       ├── help.sh
│       ├── length.sh
│       ├── letters.sh
│       ├── lowercase.sh
│       ├── numbers.sh
│       ├── uppercase.sh
│       └── test.sh
│
├── python/
│   ├── password_generator.py  ← Application entry point and all logic
│   ├── test_module.py         ← Unit tests (unittest)
│   └── readme.md              ← Python-specific notes
│
└── node/
    ├── package.json           ← npm manifest and convenience scripts
    ├── tsconfig.json          ← TypeScript compiler configuration
    ├── passwords.txt          ← Passwords appended when --save is used
    ├── readme.md              ← Node-specific notes
    └── src/
        ├── index.ts           ← CLI entry point (flag parsing, clipboard, output)
        └── utils/
            ├── createPassword.ts  ← Password generation logic
            └── savePassword.ts    ← File persistence logic
```

---

## Contributing

Contributions are welcome. If you would like to report a bug, suggest a feature, or submit a code change, please follow the standard GitHub workflow:

1. **Fork** the repository to your own GitHub account.
2. **Clone** your fork locally:

   ```bash
   git clone https://github.com/<your-username>/cli-password-generators.git
   ```

3. **Create a branch** for your work:

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make your changes.** If you are adding a feature or fixing a bug in one implementation, consider whether a parallel change is appropriate in the other two implementations to keep them consistent.
5. **Run the tests** for any implementation you modified before committing:
   - Go: `go test ./...`
   - Python: `python3 -m unittest test_module.py -v`
6. **Commit your changes** with a clear, descriptive commit message.
7. **Push** your branch to your fork:

   ```bash
   git push origin feature/your-feature-name
   ```

8. **Open a pull request** targeting the `main` branch of the original repository. Describe what you changed and why.

Please ensure that your code matches the existing style and conventions of the file you are editing. Keep pull requests focused. One logical change per PR makes review faster and easier.

---

## License

This project is licensed under the [MIT License](LICENSE).
