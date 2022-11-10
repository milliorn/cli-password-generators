const program = require("commander");

program.version("1.0.0").description("Passowrd Generator in Node.js");

/** node index generate */
program
  .command("generate")
  .action(() => {
    console.log("Generate password");
  })
  .parse();
