package main

import "fmt"

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
func main() {
	testClosure1()
}
