package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	for idx, runevalue := range s {
		fmt.Printf("%#U starts at %d\n", runevalue, idx)
	}

	fmt.Println("\nusing DecodeRuneString")

	for i, w := 0, 0; i < len(s); i += w {
		runevalue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runevalue, i)
		w = width
		examineRune(runevalue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found fee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
