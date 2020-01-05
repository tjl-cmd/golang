package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user *User
	fmt.Printf("user:%#v\n", user)
	var user01 *User = &User{}
	user01.Age = 12
	fmt.Printf("%#v\n", user01)

	var user02 *User = &User{
		Username:  "user02",
		Sex:       "男",
		Age:       123,
		AvatarUrl: "",
	}
	fmt.Printf("%v\n", user02)
	user04 := new(User)
	fmt.Printf("%v\n", user04)
}
