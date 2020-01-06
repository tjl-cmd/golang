package main

import (
	"fmt"
	"time"
)

func process(i int, ch chan bool) {
	fmt.Println("start goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	ch <- true
}
func main() {
	no := 3
	exitchan := make(chan bool, no)
	for i := 0; i < no; i++ {
		go process(i, exitchan)
	}
	for i := 0; i < no; i++ {
		<-exitchan
	}
	fmt.Println("All go routines finished executing")
}
