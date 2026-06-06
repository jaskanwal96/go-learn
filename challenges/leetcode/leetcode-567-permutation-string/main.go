package main

import (
	"fmt"
	"sort"
)

func SortString(s string) string {
	runes := []rune(s)
	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func checkInclusion(pattern string, text string) bool {
	patternLength := len(pattern)
	testPattern := SortString(pattern)
	for index := range text {
		if index <= len(text)-patternLength {
			matchingString := text[index : index+patternLength]
			fmt.Println(SortString(matchingString))
			fmt.Println(testPattern)
			if testPattern == SortString(matchingString) {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}
