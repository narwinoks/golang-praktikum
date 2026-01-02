package main

import "fmt"

func main() {
	var status string
	var duration int

	totalStorms := 0
	totalDuration := 0

	for {
		fmt.Scan(&status, &duration)

		if status == "No_Storm" {
			if duration > 0 {
				break
			}
			continue
		}

		if status == "Storm" {
			totalStorms++
			totalDuration += duration
		}
	}

	fmt.Printf("Total Storm = %d\n", totalStorms)
	fmt.Printf("Total Duration = %d minutes\n", totalDuration)
}
