package main

import "fmt"

func main() {
	var a map[string]int = map[string]int{
		"sut01": 1230,
		"stu02": 2343,
		"stu03": 2342,
	}
	fmt.Println(a)
	var key string = "stu02"
	fmt.Printf("the value of %d", a[key])
}
