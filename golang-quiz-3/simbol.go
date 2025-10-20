package main

import "fmt"

func isSpecial(c rune) bool {
	var isAlphabet, isDigit bool
	isAlphabet = (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')

	isDigit = (c >= '0' && c <= '9')

	return !isAlphabet && !isDigit
}

func main() {
	var char1, char2 rune
	var result1, result2 bool
	fmt.Scanf("%c%c", &char1, &char2)

	result1 = isSpecial(char1)

	result2 = isSpecial(char2)

	fmt.Println(result1, result2)
}
