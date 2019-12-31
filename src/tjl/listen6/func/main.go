package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func testFunc1() {
	f1 := add
	fmt.Printf("%T\n", f1)
	sum := f1(2, 43)
	fmt.Printf("sum=%d\n", sum)
}
func testFunc2() {
	f2 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", f2)
	fmt.Printf("sum=%d\n", f2(123, 234))
}
func testFunc3() {
	var i int = 0
	defer func() {
		fmt.Printf("defer i = %d \n", i)
	}()
	i = 100
	fmt.Println(i)
}
func sub(a, b int) int {
	return a - b
}

func calc(a, b int, op func(int, int) int) int {
	return op(a, b)
}
func testFunc5() {
	sum := calc(100, 200, add)
	sub := calc(200, 100, sub)
	fmt.Printf("sum=%d sub=%d", sum, sub)
}
func main() {
	//testFunc1()
	//testFunc2()
	//testFunc3()
	testFunc5()
}
