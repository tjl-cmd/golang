package main

import "fmt"

func modify(a *int) {
	*a = 100
}
func main() {
	var a = 100
	fmt.Printf("before modify:%d ,addr:%p\n", a, &a)
	modify(&a)
	fmt.Printf("after modify:%d,addr:%p\n", a, &a)
}
