package main

import "fmt"

func main() {
	nums := []int{4, 6, 2, 15, 7}
	large := nums[0]

	for count := 0; count < len(nums); count++{
		if nums[count] > large {
			large = nums[count]
		}
	}
	fmt.Println(large)
}
