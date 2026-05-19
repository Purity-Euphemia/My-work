package main

import "fmt"

func duplicate (arr []int) {
	counter := make(map[int]int)
	for _, num := range arr {
		counter[num]++
	}

	for num, count := range counter {
		if count > 1 {
			fmt.Printf("%d appears %d times\n", num, count)
		}
	}
}
