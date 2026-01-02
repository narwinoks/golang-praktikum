package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	switch {
	case number < 10:
		fmt.Println("Units")
	case number < 100:
		fmt.Println("Tens")

	case number < 1_000:
		fmt.Println("Hundreds")

	case number < 10_000:
		fmt.Println("Thousands")

	case number < 100_000:
		fmt.Println("Tens of thousands")

	case number <= 999_999:
		fmt.Println("Hundreds of thousands")

	default:
		fmt.Println("Number out of range")
	}
}
