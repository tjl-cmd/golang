package main

import "fmt"

func testIf() {
	var a int = 2
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2")
	} else {
		fmt.Println(a)
	}
}
func testSwitch() {
	var a int = 4
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("不在范围内")
	}
}
func testSwitch2() {
	var a int = 45
	switch {
	case a <= 4:
		fmt.Println("1-4")
	case a > 4 && a < 8:
		fmt.Println("5-8")
	default:
		fmt.Println("不在范围内")
	}
}
func testMulti() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
func main() {
	//testIf()
	//testSwitch2()
	testMulti()
}
