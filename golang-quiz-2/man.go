package main

import "fmt"

func main() {
	var na, nb int
	var nc float64
	var t1, t2 string
	var as bool
	fmt.Print("Masukan na dan nb: ")
	fmt.Scanln(&na, &nb)

	fmt.Println("Number 1: ", na)
	fmt.Println("Number 2: ", nb)
	t1 = "Belajar"
	t2 = "Go"
	fmt.Println("Sum na + nb =", na+nb)
	fmt.Println("Sub na - nb =", na-nb)
	fmt.Println("Mul na x nb =", na*nb)
	fmt.Println("Div na : nb =", float64(na)/float64(nb))
	fmt.Println("Res na mod nb =", na%nb)
	fmt.Println("Res na div nb =", na/nb)
	nc = float64(na) + 0.5 + 2.3
	fmt.Println(nc)
	nc = float64(nb) * 7.2 * 3.1
	fmt.Println(nc)
	t1 = t1 + " "
	fmt.Println("Teks", t1+t2)
	fmt.Println("na > nb?", na > nb)
	as = na == nb || na < nb
	fmt.Println("na is equal with nb or na is less than nb?", as)
	as = na != nb && na > 0
	fmt.Println("na is not equal with nb and na is positive?", as)
}
