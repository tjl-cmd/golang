package main

import "fmt"

func sumArray(b [10]int) int {
	sum := 0
	for _, v := range b {
		sum += v
	}
	return sum
}
func testArraySum() {
	var b [10]int
	for i := 0; i < len(b); i++ {
		b[i] = i
	}
	sum := sumArray(b)
	fmt.Printf("sum=%d", sum)
}
func TowSum(a []int, target int) {
	fmt.Println(len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[i]+a[j] == target {
				fmt.Println(i, j)
			}
		}
	}
}
func TestSum() {
	var arr []int
	arr = []int{1, 2, 3, 4, 45, 5, 6, 7, 8, 9, 10, 23, 34, 45, 3, 234, 8, 45}
	TowSum(arr, 20)
}
func main() {
	//testArraySum()
	TestSum()
}
