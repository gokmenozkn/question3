package main

import (
	"fmt"
	"sort"
)

func findMostRepeated(data []string) string {
	// Use a map to count data by frequency
	countMap := make(map[string]int)

	// Count the data
	for _, value := range data {
		countMap[value]++
	}

	// Sort numbers to find most occurrences
	var counts []int
	for _, count := range countMap {
		counts = append(counts, count)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	// Find the most frequently repeated
	for key, value := range countMap {
		if value == counts[0] {
			return key
		}
	}

	// Return an empty string if given an empty array/list or other error
	return ""
}

func main() {
	input := []string{"apple", "pie", "apple", "red", "red", "red"}
	input2 := []string{"test", "john", "john"}
	input3 := []string{""}

	output1 := findMostRepeated(input)
	output2 := findMostRepeated(input2)
	output3 := findMostRepeated(input3)

	fmt.Println("output1:", output1)
	fmt.Println("output2:", output2)
	fmt.Println("output3:", output3)
}
