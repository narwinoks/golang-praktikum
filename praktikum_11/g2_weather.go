package main

import "fmt"

func main() {
	var month string
	fmt.Scan(&month)

	switch month {
	case "December", "January", "February":
		fmt.Println("Winter")
		fmt.Println("-10 to 0 degrees C")

	case "March", "April", "May":
		fmt.Println("Spring")
		fmt.Println("5 to 15 degrees C")

	case "June", "July", "August":
		fmt.Println("Summer")
		fmt.Println("20 to 35 degrees C")

	case "September", "October", "November":
		fmt.Println("Autumn/Fall")
		fmt.Println("10 to 20 degrees C")
	default:
		fmt.Println("Invalid month name")
	}
}
