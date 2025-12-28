package main

import (
	"fmt"
	"math"
)

func main() {
	var jenisKapal string
	var posX, posY float64

	fmt.Scanln(&jenisKapal)
	fmt.Scanln(&posX)
	fmt.Scanln(&posY)
	jarak := math.Sqrt(posX*posX + posY*posY)

	const batasJarak = 5.0

	if jenisKapal == "tempur" && jarak <= batasJarak {
		fmt.Println("tembak")
	} else {
		fmt.Println("tidak tembak")
	}
}
