package main

import (
	"fmt"
)

func hello(c chan bool) {

	fmt.Println("hello goroutine")
	c <- true
}
func main() {
	var exitChan chan bool
	exitChan = make(chan bool)
	go hello(exitChan)
	fmt.Println("main thread terminate")
	//time.Sleep(time.Second)
	<-exitChan
}
