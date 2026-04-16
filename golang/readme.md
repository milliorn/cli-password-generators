# Go Password Generator

A command-line password generator written in Go. Generates a cryptographically reasonable random password from configurable character pools.

## Prerequisites

- [Go](https://go.dev/dl/) **1.22 or later**

The only external dependency is [`github.com/urfave/cli/v2`](https://github.com/urfave/cli), managed automatically by Go modules.

## Installation

```bash
cd golang
go mod download
go build -o password-generator .
```

Or run without building:

```bash
go run main.go [flags]
```

## Usage

```text
create-password [global options]
```

Running with no arguments generates an 8-character password using all character classes.

## Flags

| Flag | Short | Default | Description |
| --- | --- | --- | --- |
| `--length` | `-l` | `8` | Total length of the generated password |
| `--no-letters` | `-nle` | `false` | Exclude all letters (uppercase and lowercase) |
| `--no-lowercase` | `-nl` | `false` | Exclude lowercase letters only |
| `--no-numbers` | `-nn` | `false` | Exclude digits |
| `--no-symbols` | `-ns` | `false` | Exclude special symbols |
| `--no-uppercase` | `-nu` | `false` | Exclude uppercase letters only |
| `--version` | `-v` | n/a | Print the application version and exit |
| `--help` | `-h` | n/a | Print help text and exit |

> `--no-letters` takes priority over `--no-lowercase` and `--no-uppercase`.

## Examples

```bash
# Default: 8-character password with all character classes
./password-generator

# 12-character password without numbers
./password-generator --length 12 --no-numbers

# Numbers and symbols only
./password-generator --no-letters

# Uppercase letters and digits only
./password-generator --no-lowercase --no-symbols

# 32-character password with only digits and symbols
./password-generator --no-letters --length 32
```

## Tests

```bash
go test ./...

# Verbose output
go test -v ./...
```
