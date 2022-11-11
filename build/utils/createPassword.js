"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var random_seed_1 = __importDefault(require("random-seed"));
var alpha = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ';
var numbers = '0123456789';
var symbols = '!@#$%^&*_-+=';
function createPassword(length, hasNumbers, hasSymbols) {
    if (length === void 0) { length = 8; }
    if (hasNumbers === void 0) { hasNumbers = true; }
    if (hasSymbols === void 0) { hasSymbols = true; }
    var chars = alpha;
    hasNumbers ? (chars += numbers) : '';
    hasSymbols ? (chars += symbols) : '';
    return generatePassword(length, chars);
}
exports.default = createPassword;
var generatePassword = function (length, chars) {
    var password = '';
    /** https://github.com/skratchdot/random-seed */
    var rand = random_seed_1.default.create();
    for (var i = 0; i < length; i++) {
        var n = rand(chars.length); // generate a random number
        password += chars.charAt(n);
    }
    return password;
};
