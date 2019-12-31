package main

import "fmt"

func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func example1() {
	for i := 100; i < 1000; i++ {
		if is_shuixinahua(i) {
			fmt.Printf("[%d] is shuixinahua\n", i)
		}
	}
}
func is_shuixinahua(n int) bool {
	first := n % 10
	second := n / 10 % 10
	three := n / 100 % 10
	sum := first*first*first + second*second*second + three*three*three
	if sum == n {
		return true
	}
	return false
}
func calc(str string) (charCount, numCount, spaceCount, otherCount int) {
	utfchars := []rune(str)
	for i := 0; i < len(utfchars); i++ {
		if utfchars[i] >= 'a' && utfchars[i] <= 'z' || utfchars[i] >= 'A' && utfchars[i] <= 'Z' {
			charCount++
			continue
		}
		if utfchars[i] >= '0' && utfchars[i] <= 9 {
			numCount++
			continue
		}
		if utfchars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}
func example3() {
	var str string = "djsa   我爱艾欧尼亚  234234 '';'"
	charCount, numCount, spCount, otherCount := calc(str)
	fmt.Printf("char count:%d numcount:%d spcount:%d othercount:%d", charCount, numCount, spCount, otherCount)
}
func main() {
	//example1()
	example3()
}
