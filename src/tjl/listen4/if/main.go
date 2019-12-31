package main

import "fmt"

func testIf() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num %d not odd\n", num)
	}
}
func testIf2() {
	num := 10
	if num > 5 && num <= 10 {
		fmt.Printf("num :%d is > 5 and < 10\n", num)
	} else if num > 10 && num < 20 {
		fmt.Printf("num :%d is >10 and <20\n", num)
	} else {
		fmt.Printf("num %d is > 30\n", num)
	}
}
func testIf3() {
	if num := 10; num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num %d not odd\n", num)
	}
}
func getnum() int {
	return 10
}
func testIf4() {
	fmt.Println(getnum())
}
func main() {
	testIf4()
}
