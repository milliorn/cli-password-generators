# node-cli-password-generator

## Usage

Install dependencies

`npm install`

## Run file

`node index (options)`

## Options

| Short | Long              | Description                     |
| ----- | ----------------- | ------------------------------- |
| -l    | --length <number> | length of password (default: 8) |
| -s    | --save            | save password to passwords.txt  |
| -nn   | --no-numbers      | remove numbers                  |
| -ns   | --no-symbols      | remove symbols                  |
| -h    | --help            | display help for command        |
| -V    | --version         | Show the version                |

## Examples

```sh
node index --length=20

vsmmqiIsL%c@H^6d7nq2
```

```sh
node index -s

Password saved to passwords.txt
```

```sh
node index --length=20 -nn

Generated Password -> GCE*BJWodyZKTprsc^cY
```
