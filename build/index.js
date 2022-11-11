"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var chalk_1 = __importDefault(require("chalk"));
var clipboardy_1 = __importDefault(require("clipboardy"));
var commander_1 = require("commander");
var createPassword_js_1 = __importDefault(require("./utils/createPassword.js"));
var savePassword_js_1 = __importDefault(require("./utils/savePassword.js"));
commander_1.program.version('1.0.0').description('Password Generator');
commander_1.program
    .option('-l, --length <number>', 'length of password', '8')
    .option('-s, --save', 'save password to passwords.txt')
    .option('-nn, --no-numbers', 'remove numbers')
    .option('-ns, --no-symbols', 'remove symbols')
    .parse();
/** deconstruct options passed in */
var _a = commander_1.program.opts(), length = _a.length, save = _a.save, numbers = _a.numbers, symbols = _a.symbols;
var createdPassword = (0, createPassword_js_1.default)(length, numbers, symbols);
if (save) {
    (0, savePassword_js_1.default)(createdPassword);
}
/** copy to clipboard */
clipboardy_1.default.writeSync(createdPassword);
/** console.log(chalk.greenBright("Generated Password -> " + chalk.bold(createdPassword))) */
console.log(chalk_1.default.yellow('Password copied to clipboard'));
