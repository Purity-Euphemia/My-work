package main

import "fmt"

func main() {
	nums := []int{8, 3, 12, 1, 9}
	smallest := nums[0]

	for count := 0; count < len(nums); count++ {
		if nums[count] < smallest {
			smallest = nums[count]
		}
	}
	fmt.Println(smallest)
}
