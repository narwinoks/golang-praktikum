package main

import "fmt"

func main() {
	var color1, color2 string
	fmt.Scan(&color1, &color2)
	switch {
	case color1 == color2:
		fmt.Println(color1)
	case (color1 == "red" && color2 == "yellow") || (color1 == "yellow" && color2 == "red"):
		fmt.Println("orange")

	case (color1 == "red" && color2 == "blue") || (color1 == "blue" && color2 == "red"):
		fmt.Println("purple")

	case (color1 == "blue" && color2 == "yellow") || (color1 == "yellow" && color2 == "blue"):
		fmt.Println("green")
	default:
		fmt.Println("unknown combination")
	}
}
