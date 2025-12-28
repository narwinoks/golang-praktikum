package main

import "fmt"

func main() {
	var input1, input2, result int

	fmt.Scanln(&input1, &input2)

	for {
		result = result + input1
		input1++

		if input1 > input2 {
			break
		}
	}

	fmt.Println(result)
}
