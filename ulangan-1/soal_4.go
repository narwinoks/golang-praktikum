package main

import "fmt"

func hitungKelilingSegitiga(sisi1 float64, sisi2 float64, sisi3 float64) float64 {
	return sisi1 + sisi2 + sisi3
}
func main() {
	var sisi1, sisi2, sisi3, keliling float64
	fmt.Scanln(&sisi1, &sisi2, &sisi3)
	keliling = hitungKelilingSegitiga(sisi1, sisi2, sisi3)
	fmt.Println(keliling)
}
