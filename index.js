import chalk from "chalk";
import clipboard from "clipboardy";
import { program } from "commander";
import createPassword from "./utils/createPassword.js";
import savePassword from "./utils/savePassword.js";

program.version("1.0.0").description("Password Generator");

program
  .option("-l, --length <number>", "length of password", "8")
  .option("-s, --save", "save password to passwords.txt")
  .option("-nn, --no-numbers", "remove numbers")
  .option("-ns, --no-symbols", "remove symbols")
  .parse();

/**deconstruct options passed in */
const { length, save, numbers, symbols } = program.opts();

const createdPassword = createPassword(length, numbers, symbols);

if (save) {
  savePassword(createdPassword);
}

/** copy to clipboard */
clipboard.writeSync(createdPassword);

console.log(
  chalk.greenBright("Generated Password -> " + chalk.bold(createdPassword))
);

console.log(chalk.yellow("Password copied to clipboard"));
