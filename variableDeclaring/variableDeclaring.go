package main

import "fmt"

func main() {
	nums := [] int {1, 2, 3, 4, 5}
	for count := 0; count < len(nums); count++{
		fmt.Println(nums[count])
	}

}

