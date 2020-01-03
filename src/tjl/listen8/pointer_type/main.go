package main

import "fmt"

func testPoint1() {
	var a int32
	a = 10
	var b *int32
	b = &a
	fmt.Printf("%p,%d\n", b, *b)
	fmt.Printf("%p,%d\n", &a, a)
}
func testPoint2() {
	var a int = 200
	var b *int = &a
	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	*b = 1000
	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
	fmt.Printf("a指向的地址存储的值为:%d\n", a)
}
func modify(a *int) {
	*a = 89
}
func testPoint3() {
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("B:%d\n", b)
	fmt.Printf("p:%d\n", *p)

}
func modify_add(a *[3]int) {
	(*a)[0] = 90
}
func testPoint4() {
	var b [3]int = [3]int{23, 34, 45}
	p := &b
	modify_add(p)
	fmt.Println(b)
}
func testPoint5() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n", *a)
	var b *[]int = new([]int)
	fmt.Printf("*b=%v\n", *b)
	*b = make([]int, 5, 100)
	(*b)[0] = 100
	fmt.Printf("*b=%v\n", *b)
}
func modifyInt(a int) {
	a = 1000
}
func testPoint6() {
	var b int = 10
	modifyInt(b)
	fmt.Println(b)
}
func testPoint7() {
	var b int = 10
	var c *int = &b
	var a *int = c
	fmt.Printf("c=%d,b=%d,a=%d", *c, b, *a)
}
func main() {
	//testPoint1()
	//testPoint2()
	//testPoint3()
	//testPoint4()
	//testPoint5()
	//testPoint6()
	testPoint7()
}
