package main

import (
	"fmt"
	"sync"
	"time"
)

var rwlock sync.RWMutex
var x int
var wg sync.WaitGroup
var mutex sync.Mutex

func write() {
	for i := 0; i < 10; i++ {
		//rwlock.Lock()
		mutex.Lock()
		//fmt.Println("write lock")
		x = x + 1
		time.Sleep(time.Millisecond)
		mutex.Unlock()
		//fmt.Println("write unlock")
		//rwlock.Unlock()
	}

	wg.Done()
}
func read(i int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("wait for lock")
		//rwlock.RLock()
		mutex.Lock()
		fmt.Printf("goroutine:%d x =%d\n:", i, x)
		time.Sleep(time.Second)
		mutex.Unlock()
		//rwlock.RUnlock()
		//fmt.Println(" unlock")
	}

	wg.Done()
}
func main() {
	start := time.Now().UnixNano()
	wg.Add(1)
	go write()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("cost:", cost, "ms")
}
