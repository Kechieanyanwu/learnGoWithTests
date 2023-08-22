package main

import (
	"os"
	"regexp"
)

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// take a varying number of slices, returning a new slice containing the totals for each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) // appending each sub-sum to the main sub slice.
	}
	return sums
}

// take a varying number of slices, retun a new slice containing the totals for each slice passed in, excluding the first item of each slice
// from the various slice totals
func SumAllTails(numbersToSum ...[]int) []int {
	var tailSlice []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			tailSlice = append(tailSlice, 0)
		} else {
			tail := numbers[1:]
			tailSlice = append(tailSlice, Sum(tail))
		}
	}
	return tailSlice
}

// finds and returns the first set of consecutive numbers in a file
func FindConsecutiveNumbers(filename string) []byte {
	var digitRegexp = regexp.MustCompile("[0-9]+")
	b, _ := os.ReadFile(filename)
	consecNumSlice := digitRegexp.Find(b)
	returnSlice := make([]byte, 0, len(consecNumSlice))
	returnSlice = append(returnSlice, consecNumSlice...)
	return returnSlice
}
