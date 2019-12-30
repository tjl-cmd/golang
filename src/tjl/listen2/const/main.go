package main

import "fmt"

func main() {
	/*
		const a int = 100
	*/
	const (
		a int = 100
		b
	)
	const (
		e = iota
		f
		g
	)
	const (
		a1 = 1 << iota
		a2
		a3
	)
	fmt.Println(a, b)
	fmt.Println(e, f, g)
	fmt.Println(a1, a2, a3)
}
