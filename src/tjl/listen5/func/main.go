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
func calc_v1(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}
func calc_v2(a int, b ...int) int {
	sum := a
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}
func main() {
	//sayHello()

	//a := add(10,23)
	//fmt.Println(a)
	//a, b := calc(10, 20)
	//fmt.Println(a, b)
	sum := calc_v2(10, 20, 30, 40, 50, 60, 70)
	fmt.Printf("sum=%d\n", sum)
}
