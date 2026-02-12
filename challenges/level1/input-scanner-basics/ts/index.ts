import * as fs from "fs";
import * as readline from "readline";

// ----------------------
// ScannerMode (stdin line reader)
// ----------------------
export async function scannerMode() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  for await (const line of rl) {
    if (line.trim() === "") {
      break;
    }
    console.log(`You entered: ${line}`);
  }

  rl.close();
}

// ----------------------
// IndividualWordsScannerMode
// ----------------------
export async function individualWordsScannerMode() {
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  for await (const line of rl) {
    const [a, b, c] = line.trim().split(/\s+/);

    if (!a || !b || !c) {
      console.log("Invalid input, need 3 words");
      continue;
    }

    console.log(a, b, c);
  }

  rl.close();
}

// ----------------------
// FileMode (read file line by line)
// ----------------------
export async function fileMode() {
  const stream = fs.createReadStream("lines.txt");

  const rl = readline.createInterface({
    input: stream,
    crlfDelay: Infinity,
  });

  for await (const line of rl) {
    console.log(`File content: ${line}`);
  }

  rl.close();
}

// ----------------------
// main
// ----------------------
async function main() {
  console.log("Scanner Challenge (TypeScript)");
  console.log("-----------------------------");

  // await scannerMode();
  // await individualWordsScannerMode();
  await fileMode();
}

main();
