package main

import "fmt"

func main() {
	a := map[string]int{
		"steve": 23423,
		"stave": 234,
		"stavc": 345,
	}
	value, ok := a["stave"]
	if ok {
		fmt.Println(value)
	} else {

		fmt.Println("不存在")
	}
	for _, value := range a {
		fmt.Println(value)
	}
}
