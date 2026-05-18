package main

import "fmt"

func pointer(c *int) {
	*c = 8
}

func main() {
	c := 78
	pointer(&c)
	fmt.Println(c)


}
