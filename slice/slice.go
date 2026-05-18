package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	numbers := len(nums)
	for i := 0; i < numbers; i++{
		nums = append(nums, nums[i])
		
	}
	fmt.Println(nums)
}


