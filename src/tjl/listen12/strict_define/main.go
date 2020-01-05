package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	a := new(User)
	a.Age = 18
	a.AvatarUrl = "http://www.baidu.com"
	a.Sex = "男"
	a.Username = "user01"
	fmt.Printf("user.username = %s age = %d sex= %s avatar=%s\n", a.Username, a.Age, a.Sex, a.AvatarUrl)
	var user2 User = User{
		Username:  "user02",
		Sex:       "女",
		Age:       19,
		AvatarUrl: "http:/xxx.baidu.com",
	}
	fmt.Printf("user2=%#v\n", user2)
	user3 := User{
		Username:  "小明",
		Sex:       "男",
		Age:       20,
		AvatarUrl: "http://www.xiaoming.com",
	}
	fmt.Printf("user3=%#v\n", user3)
}
