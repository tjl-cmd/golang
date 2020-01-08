package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int32
var wg sync.WaitGroup
var mutex sync.Mutex

func Add() {
	for i := 0; i < 5000; i++ {
		//mutex.Lock()
		atomic.AddInt32(&x, 1)
		//x = x + 1
		//mutex.Unlock()
	}
	wg.Done()
}
func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go Add()
	}
	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("x:", x, "cost:", cost)
}
