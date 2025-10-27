package main

import "fmt"

func main() {
	var input1, input2 int
	fmt.Scanln(&input1, &input2)
	for {
		fmt.Println(input1)
		if input1 == input2 {
			break
		}
		input1++
	}
}
