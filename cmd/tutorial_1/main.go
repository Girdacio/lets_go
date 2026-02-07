package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, World!")

	var intNum int = 42
	var floatNum float64 = 3.14
	var str string = "Go is awesome!"
	var boolVal bool = true

	fmt.Printf("Integer: %d\n", intNum)
	fmt.Printf("Float: %.2f\n", floatNum)
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Boolean: %t\n", boolVal)

	var myString string = "Hello" + " " + "World!"
	fmt.Println(myString)
	fmt.Printf("The length of the string is: %d\n", len(myString))
	fmt.Printf("The rune count is: %d\n", utf8.RuneCountInString("γ"))

	var myRune rune = 'A'
	fmt.Printf("Rune: %c, Unicode code point: %U\n", myRune, myRune)

	const pi = 3.14159
	fmt.Printf("The value of pi is approximately: %.5f\n", pi)
}
