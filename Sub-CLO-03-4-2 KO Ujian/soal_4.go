package main

import (
	"fmt"
)

func main() {
	var temp int
	fmt.Scan(&temp)

	if temp < 0 {
		fmt.Println("Freezing weather")
	} else if temp < 10 {
		fmt.Println("Very Cold weather")
	} else if temp < 20 {
		fmt.Println("Cold weather")
	} else if temp < 30 {
		fmt.Println("Normal in Temp")
	} else if temp < 40 {
		fmt.Println("it's Hot")
	} else {
		fmt.Println("It's Very Hot")
	}
}
