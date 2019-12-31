package main

import "fmt"

func testDefer1() {
	defer fmt.Println("hello2")
	defer fmt.Println("hello1")
	defer fmt.Println("hello")
	fmt.Println("bbb")
}
func testDefer2() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("i=%d\n", i)
	}
	fmt.Printf("running\n")
	fmt.Printf("return\n")
}
func testDefer3() {
	var i int = 0
	defer fmt.Printf("i=%d\n", i)
	i = 1000
	fmt.Println(i)
}
func main() {
	//testDefer1()
	testDefer3()
}
