package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "response form sever1"
}
func server2(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "response form server2"
}
func main() {
	var output1 = make(chan string)
	var output2 = make(chan string)
	go server1(output1)
	go server2(output2)
	/*s1 := <-output1
	fmt.Println("s1:", s1)
	s2 := <-output2
	fmt.Println("s2", s2)
	*/
	select {
	case s1 := <-output1:
		fmt.Println("s1:", s1)
	case s2 := <-output2:
		fmt.Println("s2:", s2)
	default:
		fmt.Println("run default")
	}
}
