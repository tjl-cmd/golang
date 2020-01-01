package main

import "fmt"

func testSlice1() {
	var a []int
	fmt.Printf("a addr: %p", a)

}
func testSlice2() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[1:4]
	fmt.Println(b)
}
func testSlice3() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("array a :%v type of a :%T\n", a, a)
	b := a[2:5]
	b[0] = b[0] + 10
	fmt.Printf("%p\n", &a)
	fmt.Printf("%p\n", &b)
	fmt.Printf("array a :%v type of a :%T\n", a, a)

}
func testSlice5() {
	a := [...]int{1, 2, 3}
	s1 := a[:]
	s2 := a[:]
	s1[0] = 200
	s2[1] = 201
	fmt.Println(a)
	fmt.Println(a)
}
func testSlice4() {
	arr := make([]int, 10)
	arr[0] = 1
	arr[1] = 20
	fmt.Println(arr)
}
func main() {
	//testSlice2()
	//testSlice3()
	testSlice4()
	//testSlice5()
}
