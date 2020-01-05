package main

import "fmt"

func testScanf() {
	var a int
	var b string
	var c float32
	_, _ = fmt.Scanf("%d %s %f", &a, &b, &c)
	fmt.Printf("a=%d,b=%s c= %f", a, b, c)
}
func testScan() {
	var a int
	var b string
	var c float32
	fmt.Scan(&a, &b, &c)
	fmt.Printf("a=%d,b=%s c= %f", a, b, c)
}
func main() {
	//testScanf()
	testScan()
}
