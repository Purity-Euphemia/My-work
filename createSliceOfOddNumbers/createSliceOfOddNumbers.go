package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var oddNums []int

	for count := 0; count < len(nums); count++{
		if nums[count] % 2 != 0{
			oddNums = append(oddNums, nums[count])
		}
	}
	fmt.Println(oddNums)
}
