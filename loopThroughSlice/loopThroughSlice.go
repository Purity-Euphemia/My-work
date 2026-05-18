package main

import "fmt"

func main() {
	color := []string{"Red", "Blue", "Green"}

	for count := 0; count < len(color); count++ {
		fmt.Println(color[count])
	}
}
