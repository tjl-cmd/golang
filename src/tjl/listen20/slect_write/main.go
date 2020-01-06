package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write success")
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(time.Second)
	}
}
func main() {
	//select {}
	var output1 = make(chan string, 5)
	go write(output1)
	for s := range output1 {
		fmt.Println("recv", s)
		time.Sleep(time.Second)
	}
}
