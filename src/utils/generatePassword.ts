import RandomSeed from "random-seed";

export function generatePassword(length: number, chars: string): string {
  let password = "";
  /** https://github.com/skratchdot/random-seed */
  const rand = RandomSeed.create();

  for (let i = 0; i < length; i++) {
    const n = rand(chars.length); // generate a random number
    password += chars.charAt(n);
  }

  return password;
}
