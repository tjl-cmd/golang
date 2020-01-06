package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	//var s string
	//s = <-ch
	//fmt.Println(s)
	ch <- "hello"
	ch <- "world"
	ch <- "!"
	fmt.Println(len(ch))
	s1 := <-ch
	fmt.Println(s1)
	fmt.Println(<-ch)
}
