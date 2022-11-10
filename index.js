const program = require("commander");

program.version("1.0.0").description("Passowrd Generator in Node.js");

program
  .option("-l, --length <number>", "length of password", "8")
  .option("-s, --save", "save password to passwords.txt")
  .option("-nn, --no-numbers", "remove numbers")
  .option("-ns, --no-symbols", "remove symbols")
  .parse();

/**console.log(program.opts()); */

const { length, save, numbers, symbols } = program.opts();

/**console.log(numbers, symbols); */

const createdPassword = createPassword(length, numbers, symbols);

console.log(createdPassword);
