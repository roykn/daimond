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

	widest := os.Args[1]
	ascii := widest[0] // gets the ascii code
	if !(ascii >= asciiA && ascii <= asciiZ) {
		panic(helpString)
	}

	for i := asciiA; i <= int(ascii); i++ {
		fmt.Println(generateLine(fmt.Sprintf("%c", i), widest))
	}
	for i := int(ascii - 1); i >= int(asciiA); i-- {
		fmt.Println(generateLine(fmt.Sprintf("%c", i), widest))
	}
}

func getOffset(letter string, widest string) byte {
	return widest[0] - letter[0]
}

func getGap(letter string) int {
	asciiDiff := int(letter[0] - asciiA)
	return asciiDiff + (asciiDiff - 1)
}

func generateLine(letter string, widest string) string {
	gap := getGap(letter)
	offset := getOffset(letter, widest)
	line := ""

	for i := 0; i < int(offset); i++ {
		line += " "
	}

	if gap < 0 {
		return line + letter
	}

	line += letter
	for i := 0; i < gap; i++ {
		line += " "
	}
	line += letter

	return line
}
