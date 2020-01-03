package main

import "fmt"

func modify(a map[string]int) {
	a["123"] = 4678
}
func main() {
	a := make(map[string]int)
	a["lis"] = 123
	a["123"] = 234
	b := a
	fmt.Println(a)
	b["123"] = 456
	fmt.Println(a)
	modify(b)
	fmt.Println(a)
}
