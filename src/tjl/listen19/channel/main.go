package main

import "fmt"

func main() {
	var c chan int
	fmt.Printf("c=%v\n", c)
	c = make(chan int, 100)
	fmt.Printf("c=%v\n", c)
	c <- 100
	//var a int = 199
	/*

		data := <-c
		fmt.Println(data)
	*/
	<-c
}
