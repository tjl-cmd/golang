package main

import "fmt"

func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = !a
	fmt.Println(a)
	var b bool
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("left")
	}
	if a == true || b == false {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}
	fmt.Printf("%T", a)
}
func testInt() {
	var a int8
	a = 18
	fmt.Println(a)
	a = -12
	fmt.Println(a)
	var b int
	b = 123
	fmt.Println(b)
	var c float32 = 30.155
	var d = 30.1525
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("a=%d a = %x c = %f\n", a, a, c)
}
func testString() {
	//var a string = "Helloween"
	var a = "Helloween"
	fmt.Printf(a)
	b := a
	fmt.Println(b)
}
func main() {
	//testBool()
	//testInt()
	testString()
}
