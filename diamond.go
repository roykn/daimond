package main

import (
	"fmt"
	"os"
)

const helpString = "Expect exactly one capitalized letter between A and Z."
const asciiA = 65
const asciiZ = 90

func main() {
	if len(os.Args) != 2 {
		panic(helpString)
	}

	letter := os.Args[1]
	ascii := letter[0] // gets the ascii code
	if !(ascii >= asciiA && ascii <= asciiZ) {
		panic(helpString)
	}

	fmt.Println(generateLine(letter))
}

func generateLine(letter string) string {
	asciiDiff := int(letter[0] - asciiA)
	gap := asciiDiff + (asciiDiff - 1)

	if gap < 0 {
		return letter
	}

	line := ""
	line += letter
	for i := 0; i < gap; i++ {
		line += " "
	}
	line += letter

	return line
}
