# node-cli-password-generator

## Usage

Install dependencies

`npm install`

## Run file

```sh
    "tsc:build": "cd ./src && npx tsc"
```

`This command is found in package.json. Run that command an it will produce js files in dist. Once those files complete you can run the following command also found in package.json`

```sh
    "save": "node ./dist/index -s",
```

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
node ./dist/index --length=20 -s

vsmmqiIsL%c@H^6d7nq2
```

```sh
node ./dist/index -s

Password saved to passwords.txt
```

```sh
node ./dist/index --length=20 -nn -s

Generated Password -> GCE*BJWodyZKTprsc^cY
```
