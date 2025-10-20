package main

import "fmt"

func generateUniqueValue(a int, b int) bool {
	return 6*a*a*5*b+10 == 370
}
func main() {
	var a, b int
	fmt.Print("Masukan 2 angka :")
	fmt.Scanln(&a, &b)
	fmt.Println(generateUniqueValue(a, b))
}
