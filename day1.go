package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	left, right := readFileToLists("day1_input.txt")
	fmt.Printf("Distance: %d\n", sumOfDistances(left, right))
	fmt.Printf("Similarity: %d\n", similarity(left, right))
}

func similarity(left, right []int) int {
	counter := make(map[int]int)
	for _, v := range right {
		counter[v] = counter[v] + 1
	}
	var similarity int
	for _, v := range left {
		count, exists := counter[v]
		if exists {
			similarity += v * count	
		}
	}
	return similarity
}

func readFileToLists(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic("error reading file")
	}
	left := make([]int, 0, 1200)	
	right := make([]int, 0, 1200)	
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := strings.Split(scan.Text(), "   ") 
		l, e := strconv.Atoi(line[0])
		if e != nil {
			panic("not integer")
		}
		r, e := strconv.Atoi(line[1])
		if e != nil {
			panic("not integer")
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func sumOfDistances(left, right []int) int {
	if len(left) != len(right) {
		panic("lists are not equal")	
	}
	var sum int
	slices.Sort(left)
	slices.Sort(right)
	for i, v := range left {
		dist := int(math.Abs(float64(v - right[i]))) 
		sum += dist
	}
	return sum 
}