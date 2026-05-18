package main

import "fmt"

func main() {
	total := 0
	nums := []int{2, 5, 8, 11, 14, 7}

	for count := 0; count < len(nums); count++{
		if nums[count]% 2 == 0{
			total++
		}
	}
	fmt.Println(total)
}
