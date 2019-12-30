package main

import (
	"fmt"
	"strings"
	"tjl/listen2/access"
)

//布尔类型
func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = !a
	fmt.Println(a)
	var b bool
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("left")
	}
	if a == true || b == false {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}
	fmt.Printf("%T", a)
}

//int类型
func testInt() {
	var a int8
	a = 18
	fmt.Println(a)
	a = -12
	fmt.Println(a)
	var b int
	b = 123
	fmt.Println(b)
	var c float32 = 30.155
	var d = 30.1525
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("a=%d a = %x c = %f\n", a, a, c)
}

//操作字符串
func testString() {
	//var a string = "Helloween"
	var a = "Helloween"
	fmt.Printf(a)
	b := a
	fmt.Println(b)
	c := "hello world"
	fmt.Println(c)
	fmt.Printf("%v\n", c)
	d := `窗前明月光
		   疑是地上霜
		   举头望明月`
	fmt.Println(d)
	clen := len(d)
	fmt.Println(clen)
	//d = d+d
	//fmt.Println(d)
	//格式化返回字符串
	//c = fmt.Sprintf("%s%s",d,d)
	//fmt.Println(c)
	ips := "10.108.32.12;10.20.545.65"
	splice := strings.Split(ips, ";")
	fmt.Println(splice)

	//判断是否存在字符串
	result := strings.Contains(ips, "10.108.32.12")
	fmt.Println("是否包含10.108.32.12", result)
	//判断是否以某个字母或者字符开头或者结尾   HasPrefix --》开头
	str := "http://baidu.com"
	result1 := strings.HasPrefix(str, "http")
	fmt.Println("是否以http开头", result1)
	//结尾
	result2 := strings.HasSuffix(str, "com")
	fmt.Println("是否以com结尾", result2)
	//判断字符串从前往后出现的位置
	index := strings.Index(str, "baidu")
	fmt.Println("baidu第一次出现的位置是:", index)
	//判断字符串从后往前出现的位置
	index1 := strings.LastIndex(str, "baidu")
	fmt.Println("baidu出现的最后位置是：", index1)
	//join操作
	var str1 []string = []string{"10.120.123001.123.01", "10.4152.152.45.154", "12.321.524.12"}
	result3 := strings.Join(str1, ";")
	fmt.Println(result3)
}

//操作符运算
func testOperator() {
	var a int = 3
	if a == 2 {
		fmt.Printf("is right")
	} else {
		fmt.Printf("no right")
	}
	a = a + 100
	fmt.Println(a)
}
func testAccess() {
	fmt.Printf("access.a= %d", access.A)
}
func main() {
	//testBool()
	//testInt()
	//testString()
	//testOperator()
	testAccess()
}
