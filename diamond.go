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

func getOffset(letter string, widest string) byte {
	return widest[0] - letter[0]
}

func getGap(letter string) int {
	asciiDiff := int(letter[0] - asciiA)
	return asciiDiff + (asciiDiff - 1)
}

func generateLine(letter string) string {
	gap := getGap(letter)

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
