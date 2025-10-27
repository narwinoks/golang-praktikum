package main

import "fmt"

func checkBilanganSama(bilangAdik int, bilanganKaka int) bool {
	if bilangAdik == bilanganKaka {
		return true
	}
	if bilanganKaka*2 == bilangAdik {
		return true
	}
	return false
}
func main() {
	var bilangAdik, bilanganKaka int
	var isSama bool
	fmt.Scanln(&bilangAdik, &bilanganKaka)
	isSama = checkBilanganSama(bilangAdik, bilanganKaka)
	fmt.Println(isSama)
}
