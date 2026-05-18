package main

import "fmt"

func main() {
	nums := []int{5, 12, 25, 30, 18, 40}

	for count := 0; count < len(nums); count++ {
		if nums[count] > 20 {
			fmt.Println(nums[count])
		}
	}
}
