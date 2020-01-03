package main

import "fmt"

func main() {
	a := make(map[string]int)
	a["stuo1"] = 123
	a["1"] = 1234
	a["asd"] = 2423
	fmt.Println(a)
	delete(a, "stuo1")
	for key, _ := range a {
		delete(a, key)
	}
	fmt.Println(a)
}
