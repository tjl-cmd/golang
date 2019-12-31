package main

import (
	"fmt"
)

func testArray1() {
	var a [5]int
	b := [6]int{1, 2, 34, 4, 5, 567}
	c := [...]int{123, 234, 345, 45, 6567, 34, 5234}
	fmt.Println(c)
	fmt.Println(b)
	a[0] = 1234
	fmt.Println(a)
}
func testArray2() {
	a := [5]int{1, 2, 32, 4, 45}
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("a[%d]=%d\n", i, a[i])
	//}
	for index, val := range a {
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}
func testArray3() {
	a := [3][2]int{{23, 345}, {234, 456}, {234, 345}}
	for _, k := range a {
		for _, v := range k {
			fmt.Println(v)
		}
	}
}
func testArray4() {
	a := [3]int{10, 20, 30}
	b := a
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
func testAdd() {
	a := [3]int{12, 23, 45}
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
func main() {
	//testArray1()
	//testArray2()
	//testArray3()
	//testArray4()
	testAdd()
}
