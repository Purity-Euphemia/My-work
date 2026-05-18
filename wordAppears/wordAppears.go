package main

import (
	"fmt"
	"strings"
)

func appears (a string) map[string]int {
	i := strings.Fields(a)
	counter := make(map[string]int)
	for _, i := range i {
		counter[i]++
	}
	return counter
} 

func main() {
	c := "sweet"
	fmt.Println(appears(c))

}
