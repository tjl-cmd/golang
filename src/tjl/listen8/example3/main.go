package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr   = "0123456789"
	CharStr  = "abcdefghijklimnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpaceStr = "!@#$%^&*"
)

func parseArgs() {
	flag.IntVar(&length, "1", 16, "-1 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成的字符集，
				num只使用数字[0-9],
				char只使用英文字母[a-zA-Z],
				mix:使用数字和字母，
				advance：使用数字、字母以及特殊字符串`)
	flag.Parse()
}
func generatePassword() string {
	var password []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpaceStr)
	} else {
		sourceStr = NumStr
	}
	fmt.Println("source", sourceStr)

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}
	return string(password)
}
func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	fmt.Printf("length:%d,charset:%s\n", length, charset)
	passwd := generatePassword()
	fmt.Println(passwd)
}
