import chalk from "chalk";
import clipboard from "clipboardy";
import { program } from "commander";
import createPassword from "./utils/createPassword.js";
import savePassword from "./utils/savePassword.js";

program.version("1.0.0").description("Password Generator");

program
  .option("-l, --length <number>", "Length of password", "8")
  .option("-s, --save", "Save password to passwords.txt")
  .option("-n, --no-numbers", "Exclude numbers")
  .option("-y, --no-symbols", "Exclude symbols")
  .parse();

/** deconstruct options passed in */
const { length, save, numbers, symbols } = program.opts();

const createdPassword = createPassword(length, numbers, symbols);

if (save) {
  savePassword(createdPassword);
}

/** copy to clipboard */
clipboard.writeSync(createdPassword);

/** console.log(chalk.greenBright("Generated Password -> " + chalk.bold(createdPassword))) */

console.log(chalk.yellow("Password copied to clipboard"));
