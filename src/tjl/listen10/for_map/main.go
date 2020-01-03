package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a map[string]int = make(map[string]int, 1024)
	for i := 0; i < 1024; i++ {
		value := rand.Intn(1000)
		key := fmt.Sprintf("stu%d", i)
		a[key] = value
	}
	for key, value := range a {
		fmt.Printf("map[%s]=%d\n", key, value)
	}
}
