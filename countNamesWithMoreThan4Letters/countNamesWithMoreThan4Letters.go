package main

import "fmt"

func main() {
	total := 0
	names := []string{"Ada", "Chioma", "John", "Emaka", "Tobi"}

	for count := 0; count < len(names); count++{
		if len(names[count]) > 4{
			total++
		}
	}
	fmt.Println(total)
}