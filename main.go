package main // import "cpl.li/go/alpha"

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage:\n    alpha \"single sentence\" or multiple words\n")
		fmt.Printf("\nlicense:\n    DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE")
		os.Exit(1)
	}

	for _, str := range os.Args[1:] {
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			continue
		}

		out, count := stringToAlphas(str)
		if count == 0 {
			continue
		}

		fmt.Println(out)
	}
}

var dict = map[rune]string{
	'A': "Alpha",
	'B': "Bravo",
	'C': "Charlie",
	'D': "Delta",
	'E': "Echo",
	'F': "Foxtrot",
	'G': "Golf",
	'H': "Hotel",
	'I': "India",
	'J': "Juliett",
	'K': "Kilo",
	'L': "Lima",
	'M': "Mike",
	'N': "November",
	'O': "Oscar",
	'P': "Papa",
	'Q': "Quebec",
	'R': "Romeo",
	'S': "Sierra",
	'T': "Tango",
	'U': "Uniform",
	'V': "Victor",
	'W': "Whiskey",
	'X': "X-Ray",
	'Y': "Yankee",
	'Z': "Zulu",
	'0': "Zero",
	'1': "One",
	'2': "Two",
	'3': "Three",
	'4': "Four",
	'5': "Five",
	'6': "Six",
	'7': "Seven",
	'8': "Eight",
	'9': "Nine",
}

func stringToAlphas(str string) ([]string, int) {
	alphas := make([]string, 0)
	count := 0

	for _, char := range str {
		alpha, ok := dict[unicode.ToUpper(rune(char))]
		if !ok {
			alphas = append(alphas, "_")
			continue
		}

		count++
		alphas = append(alphas, alpha)
	}

	return alphas, count
}
