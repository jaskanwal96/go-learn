# Scanner Challenge (TypeScript)

This challenge introduces you to reading input from the console using Node.js `readline` module.

## Goal
The program should prompt the user for input, read a line of text, and then print it back to the console.

## How to Run
1. Navigate to this directory:
   ```bash
   cd challenges/level1/scanner/ts
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Run the program:
   ```bash
   npm start
   ```

## Key Concepts
- `readline.createInterface`: Creates an interface for reading from a stream (like `process.stdin`).
- `rl.question(query, callback)`: Displays the query and invokes the callback with the user's input.
- `rl.close()`: Closes the readline interface.
