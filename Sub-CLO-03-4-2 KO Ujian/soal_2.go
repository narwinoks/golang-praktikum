package main

import "fmt"

func main() {
	var p, v, ga, fs, ci, tc int

	fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

	total := p + v + ga + fs + ci + tc

	rataRata := float64(total) / 6.0

	if rataRata >= 3.33 {
		fmt.Println("mendapat sertifikat")
	} else {
		fmt.Println("tidak mendapat sertifikat")
	}
}
