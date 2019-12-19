package main

import "fmt"

func add(a int, b int) (int, string) {
	c := a + b
	return c, "successfully added"
}

func main() {
	sum, message := add(2, 1)
	fmt.Println(message)
	fmt.Println(sum)
}
