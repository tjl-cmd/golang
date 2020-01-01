package main

import "fmt"

func testMake2() {
	var a []int
	a = make([]int, 5, 10)
	a = append(a, 10)
	fmt.Printf("a=%v\n", a)
	fmt.Println(cap(a))
	b := a
	b = append(b, 1203)
	fmt.Println(a)
	fmt.Println(b)
}
func testMake3() {
	a := [...]string{"a", "b", "c", "d", "e"}
	b := a[1:3]
	fmt.Printf("b:%v,len%d,cap:%d\n", b, len(b), cap(b))
}
func main() {
	//testMake2()
	testMake3()
}
