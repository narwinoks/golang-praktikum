package main

import "fmt"

func main() {
	var n int
	var target1, target2 string

	fmt.Scan(&n, &target1, &target2)

	var currentStore string
	var count1, count2 int = 0, 0
	var i int = 0

	for i < n {
		fmt.Scan(&currentStore)

		if currentStore == target1 {
			count1++
		} else if currentStore == target2 {
			count2++
		}

		i++
	}

	if count1 > count2 {
		fmt.Println(target1, target2)
	} else if count2 > count1 {
		fmt.Println(target2, target1)
	} else {
		if target1 < target2 {
			fmt.Println(target1, target2)
		} else {
			fmt.Println(target2, target1)
		}
	}
}
