package main

import "fmt"

func testString() {
	var str = "hello"

	fmt.Printf("str[0]=%c\n", str[0])
	//for i :=0; i<len(str) ;i++  {
	//	fmt.Println(str[i])
	//}
	for index, val := range str {
		fmt.Printf("str[%d]=%c\n", index, val)
	}
	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = 'o'
	str = string(byteSlice)
	fmt.Println(str)
	fmt.Printf("len(str)=%d\n", len(str))
	//一个中文字符占3个字节
	str = "hello 少林之巅"
	fmt.Printf("len(str)=%d\n", len(str))
	//定义一个utf8字符 rune
	var b rune = '中'
	fmt.Printf("b=%c\n", b)
	//定义一个utf8切片
	var runeSlice []rune
	//将字符串转成切片
	runeSlice = []rune(str)
	//求字符串的个数长度
	fmt.Printf("str 长度:%d\n", len(runeSlice))
}

//对英文字母进行逆序
func testReverseStringV1() {
	var str = "hello"
	fmt.Println(len(str))
	var bytes []byte = []byte(str)
	for i := 0; i < len(bytes)/2; i++ {
		temp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = temp
	}
	str = string(bytes)
	fmt.Println(str)
}

//对中文进行逆序
func testReverseStringV2() {
	var str = "hello中文 "
	//fmt.Println(len(str))
	var r []rune = []rune(str)
	//fmt.Println(len(r))
	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
		//fmt.Println(i)
		//fmt.Println(len(r)-i-1)
		//fmt.Printf("%c",r[i])
		//r[len(r)-i-1],r[i] = r[i],r[len(r)-i-1]
	}
	str = string(r)
	fmt.Println(str)
}

//判断一个字符串是否是回文
func testHuiwen() {
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)
	for i := 0; i < len(r)/2; i++ {
		r[len(r)-i-1], r[i] = r[i], r[len(r)-1-i]
	}
	str2 := string(r)
	if str2 == str {
		fmt.Println("是回文")
	} else {
		fmt.Println("不是回文")
	}
}

//字符串操作原理
func main() {
	//testString()
	//testReverseStringV1()
	//testReverseStringV2()
	testHuiwen()
}
