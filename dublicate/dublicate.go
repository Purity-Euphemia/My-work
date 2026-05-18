package main

import "fmt"

func duplicate(nums []int) []int {
	seen := make(map[int]bool)
	duplicates := []int{}

	for _, num := range nums {
		if seen[num] {
			duplicates = append(duplicates, num)
		} else {
			seen[num] = true
		}
	}
	return duplicates
}

func main() {
	nums := []int{1,2,3,2,4,5,1}
	fmt.Println(duplicate(nums))
}
