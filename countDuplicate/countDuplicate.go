package main

import "fmt"

func countDuplicate(nums []int) {
	counts := map[int]int{}

	for count := 0; count < len(nums); count++{
		counts[nums[count]]++
	}
	for key, value := range counts {
		if value > 1 {
			fmt.Println(key, "appears", value, "times")
		}
	}
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	countDuplicate(nums)
}
