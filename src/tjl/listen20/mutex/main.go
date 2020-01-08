package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup

//互斥锁
var mutex sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("x:", x)
}
