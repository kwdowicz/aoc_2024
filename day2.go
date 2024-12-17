package main

import (
	"bufio"
	_ "fmt"
	"math"
	"strconv"
	"strings"
	"os"
)

func readReportFile(fileName string) [][]int {
	reports := make([][]int, 0, 1000)
	file, err := os.Open(fileName)
	if err != nil {
		panic("cant read file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dps := strings.Split(line, " ")
		report := make([]int, 0, 100)
		for _, dp := range dps {
			i, e := strconv.Atoi(dp)
			if e != nil {
				panic("cant convert to integer")
			}
			report = append(report, i)
		}
		reports = append(reports, report)
	}
	return reports
}

func reviewReport(r []int) bool {
	var trend int = 0 
	for i := 0; i < len(r)-1; i++ {
		this := r[i]
		next := r[i+1]

		if this == next {
			return false
		}

		if int(math.Abs(float64(this - next))) > 3 {
			return false
		}

		if this > next {
			if trend == 1 {
				return false
			}
			trend = -1
		} else {
			if trend == -1 {
				return false
			}
			trend = 1
		}

		if this > next && trend == 1 {
			return false
		}

		if this < next && trend == -1 {
			return false
		}

	}
	return true
}