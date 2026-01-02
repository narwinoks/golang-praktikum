package main

import "fmt"

func main() {
	var n int
	var c1, c2, c3 float64
	var cntA, cntAB, cntB, cntBC, cntC, cntFail int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&c1, &c2, &c3)

		totalScore := (c1 * 0.35) + (c2 * 0.35) + (c3 * 0.30)
		if totalScore > 85 {
			cntA++
		} else if totalScore > 75 {
			cntAB++
		} else if totalScore > 65 {
			cntB++
		} else if totalScore > 60 {
			cntBC++
		} else if totalScore > 50 {
			cntC++
		} else {
			cntFail++
		}
	}
	fmt.Printf("Grade A: %.2f %%\n", float64(cntA)/float64(n)*100)
	fmt.Printf("Grade AB: %.2f %%\n", float64(cntAB)/float64(n)*100)
	fmt.Printf("Grade B: %.2f %%\n", float64(cntB)/float64(n)*100)
	fmt.Printf("Grade BC: %.2f %%\n", float64(cntBC)/float64(n)*100)
	fmt.Printf("Grade C: %.2f %%\n", float64(cntC)/float64(n)*100)
	fmt.Printf("Not Passed: %.2f %%\n", float64(cntFail)/float64(n)*100)
}
