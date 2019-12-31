package main

import (
	"fmt"
	"strings"
	"time"
)

func Add() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}
func testClosure1() {
	f := Add()
	ret := f(1)
	fmt.Printf("ret =%d", ret)
	ret = f(20)
	fmt.Printf("ret =%d", ret)
	ret = f(100)
	fmt.Printf("ret =%d", ret)

}
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix

		}
		return name
	}
}
func testClosure3() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
func testClosure4() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
}
func testClosure5() {
	for i := 0; i < 5; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second)
}
func main() {
	//testClosure1()
	//testClosure3()
	//testClosure4()
	testClosure5()
}
