package main

import "fmt"

func hasilAIUEO(input string) bool {
	if input == "a" {
		return true
	}
	if input == "i" {
		return true
	}

	if input == "u" {
		return true
	}
	if input == "e" {
		return true
	}
	if input == "0" {
		return true
	}

	return false
}
func main() {
	var input string
	var hasil bool
	fmt.Scanln(&input)
	hasil = hasilAIUEO(input)
	fmt.Printf("%t", hasil)

}
