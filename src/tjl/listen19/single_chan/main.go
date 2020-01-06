package main

import "fmt"

//单向只写
func sedData(sendch chan<- int) {
	sendch <- 10
}

//单向只读
func readData(sendch <-chan int) {
	data := <-sendch
	fmt.Println(data)
}
func main() {
	chal := make(chan int)
	go sedData(chal)
	readData(chal)
}
