package main

import (
	"fmt"
	"tjl/listen11/calc"
)

var a int = 10

func init() {
	fmt.Println(a)
	fmt.Println("init func in  main")
}
func main() {
	var sum = calc.Add(2, 5)
	fmt.Println(sum)
}
