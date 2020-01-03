package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := make([]map[string]int, 5, 16)
	for index, val := range a {
		fmt.Printf("slice[%d]=%v\n", index, val)
	}
	a[0] = make(map[string]int, 16)
	//fmt.Println(a)
	a[0]["stu01"] = 1231
	a[0]["stu02"] = 345
	a[0]["stu03"] = 556
	a[0]["stu04"] = 1231
	a[0]["stu05"] = 345
	a[0]["stu06"] = 556
	fmt.Println(a)
	for _, val := range a {
		//fmt.Printf("slice[%d]=%v\n", index, val)
		for key, vale := range val {
			fmt.Println(key, vale)
		}
	}
}
