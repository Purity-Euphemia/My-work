package main

import "fmt"

func main() {
	fruit := []string{"mongo", "stew", "coco"}
	for count := 0; count < len(fruit); count++{
		fmt.Println(fruit[count])
	} 
	
}
