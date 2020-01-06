package main

import (
	"fmt"
	"time"
)

var i int

func calc() {
	for {
		i++
		fmt.Println(i)
	}
}
func main() {
	/*
		cpu := runtime.NumCPU()
		fmt.Println("cpu:", cpu)
		runtime.GOMAXPROCS(cpu)

	*/
	for i := 0; i < 10; i++ {
		go calc()
	}
	time.Sleep(time.Hour)
}
