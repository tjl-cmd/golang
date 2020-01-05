package main

import "fmt"

type User struct {
	Username string
	Sex      string
	int
	string
}

func main() {
	var user User
	user.Username = "user01"
	user.Sex = "男"
	user.int = 123
	user.string = "你好"
	fmt.Printf("%v\n", user)
}
