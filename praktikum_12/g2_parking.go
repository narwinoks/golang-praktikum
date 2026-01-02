package main

import "fmt"

func main() {
	var hIn, mIn, hOut, mOut int
	totalMinutes := 0

	for {
		fmt.Scan(&hIn, &mIn, &hOut, &mOut)

		switch {
		case hIn == 0 && mIn == 0 && hOut == 0 && mOut == 0:
			goto Calculate
		}

		start := hIn*60 + mIn
		end := hOut*60 + mOut
		duration := end - start

		totalMinutes += duration
	}

Calculate:
	hours := totalMinutes / 60
	remainder := totalMinutes % 60

	switch {
	case remainder > 0:
		hours++
	}

	cost := hours * 3000
	fmt.Printf("Rp %d\n", cost)
}
