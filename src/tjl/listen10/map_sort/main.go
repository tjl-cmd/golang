package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int, 1024)

	for i := 0; i < 128; i++ {
		key := fmt.Sprintf("sdu%d", i)
		value := rand.Intn(1000)
		a[key] = value
	}
	var keys []string = make([]string, 0, 128)
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		fmt.Printf("key:%s,val:%d\n", value, a[value])
	}
}
