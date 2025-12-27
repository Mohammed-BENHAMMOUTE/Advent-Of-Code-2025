import { promises as fsPromise } from 'fs';

async function solve() {
  let current_count = 50;
  let num_zeroes = 0;
  const MOD = 100;

  const content = await fsPromise.readFile("./input.txt", 'utf8');
  const lines = content.split('\n').map(line => line.trim()).filter(Boolean);

  for (const line of lines) {
    const action = line.charAt(0);
    const step = Number(line.substring(1));

    if (action === 'L') {
      current_count = (((current_count - step) % MOD) + MOD) % MOD;
    } else if (action === 'R') {
      current_count = (current_count + step) % MOD;
    }
    console.log(`Line: ${line} | Step: ${step} | Count: ${current_count}`);
    if (current_count === 0) {
      num_zeroes += 1;
    }
  }
  return num_zeroes;
}
console.log(await solve());
