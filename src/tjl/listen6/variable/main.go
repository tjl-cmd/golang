package main

import "fmt"

var a int = 100

func testGlobalVariable() {
	fmt.Printf("a=%d\n", a)
}
func testLocalVariable() {
	var b int = 101
	fmt.Println(b)
}
func main() {
	fmt.Println(a)
	testGlobalVariable()
	testLocalVariable()
}
