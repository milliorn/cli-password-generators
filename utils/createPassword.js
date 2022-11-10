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

  for (let i = 0; i < length; i++) {
    password += chars.charAt(Math.random() * chars.length);
  }

  return password;
};
