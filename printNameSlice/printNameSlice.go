package main

import "fmt"

func main() {
	names := []string{"Ada", "John", "Emaka", "Chi"}

	for count := 0; count < len(names); count++{
		if len(names[count]) > 3 {
			fmt.Println(names[count])
		}
	}
}
