package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for count := 0; count < len(nums); count++{
		if nums[count] % 2 != 0 {
			fmt.Println(nums[count])
		}
	}
}
