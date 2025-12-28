package main

import "fmt"

func main() {
	var banyakLoop, input1 int
	var result int = 1

	fmt.Scanln(&banyakLoop)
	for i := 0; i < banyakLoop; i++ {
		fmt.Scanln(input1)
		result = result * input1
	}
	fmt.Println(result)
}
