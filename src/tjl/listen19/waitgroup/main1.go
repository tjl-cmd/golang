package main

import (
	"fmt"
	"sync"
	"time"
)

func process1(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goruoutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	wg.Wait()
	fmt.Println("wait return")
	for i := 0; i < no; i++ {
		wg.Add(i)
		go process1(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
