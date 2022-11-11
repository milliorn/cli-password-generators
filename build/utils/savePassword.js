"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
var chalk_1 = __importDefault(require("chalk"));
var fs_1 = require("fs");
var os_1 = require("os");
var path_1 = __importStar(require("path"));
/** https://stackoverflow.com/a/68163774/11986604 */
var __dirname = path_1.default.resolve();
function savePassword(password) {
    (0, fs_1.open)((0, path_1.join)(__dirname, './', 'passwords.txt'), 'a', 666, function (e, id) {
        (0, fs_1.write)(id, password + os_1.EOL, null, 'utf-8', function () {
            (0, fs_1.close)(id, function () {
                console.log(chalk_1.default.green('Password saved to passwords.txt'));
            });
        });
    });
}
exports.default = savePassword;
