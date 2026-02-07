package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Println("Indexed character (byte):", indexed)
	fmt.Printf("%v, %T", indexed, indexed)

	for i, v := range myString {
		fmt.Printf("Index: %d, ASCII: %d, Character: %c\n", i, v, v)
	}

	var myRune rune = 'é'
	fmt.Printf("Rune: %c, Unicode code point: %U, value: %v\n", myRune, myRune, myRune)

	var strSlice = []string{"a", "b", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("Concatenated string: %v\n", catStr)
}
