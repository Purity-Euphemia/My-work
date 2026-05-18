package main

import "fmt"

func main() {
	names := []string{"Ada", "John", "Emake", "Chioma"}

	for count := 0; count < len(names); count++{
		if len(names[count]) < 5 {
			fmt.Println(names[count])
		}
	}
}
