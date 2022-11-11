import RandomSeed from "random-seed";

const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
const numbers = "0123456789";
const symbols = "!@#$%^&*_-+=";

export default function createPassword(
  length = 8,
  hasNumbers = true,
  hasSymbols = true
) {
  let chars = alpha;
  hasNumbers ? (chars += numbers) : "";
  hasSymbols ? (chars += symbols) : "";
  return generatePassword(length, chars);
}

const generatePassword = (length, chars) => {
  let password = "";
  /**https://github.com/skratchdot/random-seed */
  const rand = RandomSeed.create();

  for (let i = 0; i < length; i++) {
    const n = rand(chars.length); // generate a random number
    password += chars.charAt(n);
  }

  return password;
};
