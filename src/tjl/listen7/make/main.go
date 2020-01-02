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
func testCap2() {
	a := [...]string{"a", "b", "c", "d", "e"}
	b := a[1:3]
	c := b[2:cap(b)]
	fmt.Println(c)
}
func testSlice() {
	var a []int
	fmt.Printf("%p len(%d),cap(%d)\n", a, len(a), cap(a))
	if a == nil {
		fmt.Println("a in nil")
	} else {
		fmt.Println("a not nil")
	}
	a = append(a, 100)
	a = append(a, 200)
	fmt.Println(a)
	fmt.Printf("%p len(%d),cap(%d)\n", a, len(a), cap(a))

}
func testAppend() {
	var a []int = []int{1, 2, 3, 45}
	var b []int = []int{4, 5, 6, 7}
	a = append(a, 23, 456, 67, 8)
	fmt.Printf("a=%v\n", a)
	a = append(a, b...)
	fmt.Printf("a=%v\n", a)
}
func sumArray(a []int) int {
	var sum int = 0
	for _, v := range a {
		sum = sum + v
	}
	return sum
}
func testSumArray() {
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 2, 3, 34, 45, 567, 678, 78, 3, 53, 234}
	sum := sumArray(a[:])
	fmt.Printf("sum:%d", sum)
}
func testCopy() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{4, 5, 6}
	copy(a, b)
	fmt.Println(a)
}
func main() {
	//testMake2()
	//testMake3()
	//testCap2()
	//testSlice()
	//testAppend()
	//testSumArray()
	testCopy()
}
