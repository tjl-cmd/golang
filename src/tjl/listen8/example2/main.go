package main

import (
	"fmt"
	"sort"
)

//使用sort对数据进行排序

func main() {
	var a [5]int = [5]int{5, 3, 2, 4, 1}
	sort.Ints(a[:])
	fmt.Println(a)
	var b []string = []string{"ab", "ac", "ae", "be", "bd"}
	sort.Strings(b[:])
	fmt.Println(b)
	var c []float64 = []float64{23.345, 45.456, 456.34}
	sort.Float64s(c[:])
	fmt.Println(c)
}
