package main

import (
	"fmt"
	"tjl/listen12/user"
)

func main() {
	/*var u user.User
	u.Age = 19
	u.Username = "小明"
	u.Sex = "男"
	fmt.Println(u)*/
	u := user.NewUser("user01", "女", 19, "xx.jpg")
	fmt.Printf("user=%#v\n", u)
}
