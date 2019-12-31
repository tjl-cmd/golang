package main

import "fmt"

func sayHello() {
	fmt.Println("hello world")
}
func add(a, b int) int {
	return a + b
}
func calc(a, b int) (int, int) {
	sum := a + b
	sub := a * b
	return sum, sub
}
func main() {
	//sayHello()

	//a := add(10,23)
	//fmt.Println(a)
	a, b := calc(10, 20)
	fmt.Println(a, b)
}
