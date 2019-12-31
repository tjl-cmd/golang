package main

import (
	"fmt"
	"tjl/listen6/calc"
)

func main() {
	var s1 int = 100
	var s2 int = 200
	sum := calc.Add(s1, s2)
	fmt.Println(sum)
	num := calc.A
	fmt.Println(num)
}
