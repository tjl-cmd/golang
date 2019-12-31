package main

import "fmt"

func testFor() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
}
func testFor1() {
	var i int
	for i = 1; i < 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}
func testFor2() {
	var i int
	for i = 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(i)
}
func testFor3() {
	i := 0
	for i <= 10 {
		if i == 4 {
			continue
		}
		fmt.Println(i)
		i += 2

	}
	//fmt.Println(i)
}
func testFor4() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 2
	}
}
func testFor5() {
	var a int
	var b int
	for a, b = 10, 1; a <= 19 && b < 10; a, b = a+1, b+1 {
		fmt.Printf("a = %d b = %d \n", a, b)
	}
}
func main() {
	testFor5()
}
