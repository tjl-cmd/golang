package main

import "fmt"

//插入排序
func insert_sort(a [8]int) [8]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}

	return a
}

//选择排序
func select_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

//冒泡排序
func bubble_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
func main() {
	var i [8]int = [8]int{1, 2, 4, 5, 7, 8, 3, 6}
	b := bubble_sort(i)
	fmt.Println(b)
	//j := insert_sort(i)
	//k := select_sort(i)
	//fmt.Println(k)
	//fmt.Println(i)
	//fmt.Println(j)
}
