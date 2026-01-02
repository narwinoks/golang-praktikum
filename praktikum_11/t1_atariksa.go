package main

import "fmt"

func main() {
	var gravity float64
	var temp int
	var atmos, isGravBad, isTempBad, isAtmosBad bool
	fmt.Scan(&gravity, &temp, &atmos)

	isGravBad = gravity < 0.5 || gravity > 1.5
	isTempBad = temp < -50 || temp > 50
	isAtmosBad = !atmos

	switch {
	case !isGravBad && !isTempBad && !isAtmosBad:
		fmt.Println("Landing possible: planet conditions are safe for landing.")
	case isGravBad && isTempBad && isAtmosBad:
		fmt.Println("Landing not possible: planet's gravity, temperature, and atmosphere are not safe.")

	case isGravBad && isTempBad:
		fmt.Println("Landing not possible: planet's gravity and temperature are not safe.")

	case isGravBad && isAtmosBad:
		fmt.Println("Landing not possible: planet's gravity and atmosphere are not safe.")

	case isTempBad && isAtmosBad:
		fmt.Println("Landing not possible: planet's temperature and atmosphere are not safe.")

	case isGravBad:
		fmt.Println("Landing not possible: planet's gravity is not safe.")

	case isTempBad:
		fmt.Println("Landing not possible: planet's temperature is not safe.")

	case isAtmosBad:
		fmt.Println("Landing not possible: planet's atmosphere is not safe.")
	}
}
