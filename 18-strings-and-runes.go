package main

import (
	"fmt"
	"unicode/utf8"
)

/*
 * A Go string is a read-only slice of bytes ([]byte), a container of text encoded in UTF-8
 * In other languages, strings are made of characters
 * In Go, a character is called a rune - an integer that represents a Unicode code point
 */

func main() {
	const s = "สวัสดี" // string literals are UTF-8 encoded text

	fmt.Println("len:", len(s)) // length of raw bytes stored within (prints 18)

	for i := 0; i < len(s); i++ { // indexing into string produces raw byte values at each index
		fmt.Printf("%x ", s[i]) // generates hex values of all the bytes that constitute the code points in s
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s)) // utf8 package has utilities for runes

	for idx, runeValue := range s { // range loop handles strings specially
		fmt.Printf("%#U starts at %d\n", runeValue, idx) // decodes each rune along with offset in the string
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue) // can pass runes as arguments to functions
	}
}

func examineRune(r rune) {
	if r == 't' { // single quotes = rune literal, can compare a rune to a rune literal directly
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
