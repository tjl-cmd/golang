package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	sum := add(2, 6)
	fmt.Println(sum)
}
