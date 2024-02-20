import chalk from "chalk";
import { close, open, write } from "fs";
import { EOL } from "os";
import path, { join } from "path";

/** https://stackoverflow.com/a/68163774/11986604 */
const __dirname = path.resolve();

export default function savePassword(password: string): void {
  open(join(__dirname, "./", "passwords.txt"), "a", 666, (e, id) => {
    write(id, password + EOL, null, "utf-8", () => {
      close(id, () => {
        console.log(chalk.green("Password saved to passwords.txt"));
      });
    });
  });
}
