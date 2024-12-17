package main

import (
	"fmt"
)

func main() {
	/* day 1
	left, right := readFileToLists("day1_input.txt")
	fmt.Printf("Distance: %d\n", sumOfDistances(left, right))
	fmt.Printf("Similarity: %d\n", similarity(left, right))
	*/
	var safe, unsafe int
	reports := readReportFile("day2_input.txt")
	for _, r := range reports {
		if reviewReport(r) {
			safe++
		} else {
			unsafe++
		}
	}
	fmt.Printf("safe: %d, unsafe: %d\n", safe, unsafe)
}