package main

import "fmt"

func main() {
	var a map[string]int
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a=%v\n", a)
		a["stu01"] = 1023
		a["stu02"] = 1023
		a["stu03"] = 1023
		fmt.Println(a["stu01"])
	}
	fmt.Printf("a:%v", a)
}
