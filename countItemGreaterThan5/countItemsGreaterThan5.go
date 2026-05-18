package main

import "fmt"

func main() {
	nums := []int{1, 6, 3, 8, 2, 10}
	total := 0

	for count := 1; count < len(nums); count++ {
		if nums[count] > 5 {
			total++
		}
	}
	fmt.Println(total)
}
