package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ScannerMode() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			// still print the last line if any
			if len(line) > 0 {
				fmt.Printf("You entered: %s", line)
			}
			break
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		fmt.Printf("You entered: %s\n", line)
	}
}

func IndividualWordsScannerMode() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var a, b, c string
		fmt.Sscanf(scanner.Text(), "%s %s %s", &a, &b, &c)
		fmt.Println(a, b, c)
	}
}

func FileMode() {
	file, err := os.Open("lines.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadString('\n')
		if err == io.EOF {
			// still print the last line if any
			if len(line) > 0 {
				fmt.Printf("File content: %s", line)
			}
			break
		}
		fmt.Printf("File content: %s\n", line)
	}
}

func main() {
	fmt.Println("Scanner Challenge (Go)")
	fmt.Println("----------------------")
	// ScannerMode()
	IndividualWordsScannerMode()
}
