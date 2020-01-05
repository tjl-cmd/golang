package main

import (
	"flag"
	"fmt"
)

var recusive bool
var test string
var level int

func init() {

	flag.BoolVar(&recusive, "r", false, "resuive.xxx")
	flag.StringVar(&test, "t", "default string", "string option")
	flag.IntVar(&level, "1", 1, "level of xxxx")
	flag.Parse()
}

func main() {
	fmt.Printf("recusive:%v\n", recusive)
	fmt.Printf("test:%v\n", test)
	fmt.Printf("level:%v\n", level)

}
