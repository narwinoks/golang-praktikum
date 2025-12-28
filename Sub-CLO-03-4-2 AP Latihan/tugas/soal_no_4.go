package main

import (
	"fmt"
)

func main() {
	var isHariKerja, mauNonton, isGenreAksi bool
	var jamNonton int

	fmt.Scanln(&isHariKerja)
	fmt.Scanln(&jamNonton)
	fmt.Scanln(&isGenreAksi)
	mauNonton = isHariKerja && (jamNonton > 19) && isGenreAksi
	fmt.Println(mauNonton)
}
