package main

import "fmt"

/*
type User struct {
	Username string
	Sex      string
	address  Address
}*/
type User struct {
	Username string
	Sex      string
	Address
}
type User01 struct {
	Username string
	Sex      string
	City     string
	*Address
	*email
}
type Address struct {
	Province   string
	City       string
	CreateTime string
}
type email struct {
	account    string
	CreateTime string
}

func test1() {
	var user User
	user.Username = "user01"
	user.Sex = "man"
	//第一种方式
	user.Address = Address{
		Province: "bj",
		City:     "bj",
	}
	//第二种方式
	user.Province = "bj1"
	user.City = "bj1"
	fmt.Printf("user= %#v\n addr = %#v\n", user, user.Address)
}
func test02() {
	var user User01
	user.Username = "user01"
	user.City = "bj"
	user.Address = new(Address)
	user.email = new(email)

	user.Address.City = "tj"
	fmt.Printf("user=%#v\n city of address:%s\n", user, user.Address.City)
	user.Address.CreateTime = "0001"
	user.email.CreateTime = "0002"
	fmt.Printf("user=%#v\n createTime:%s email:%s\n", user, user.Address.CreateTime, user.email.CreateTime)
}
func main() {
	test02()
	/*user := &User{
		Username: "user01",
		Sex:      "man",
		address: Address{
			Province: "beijing",
			City:     "beijing",
		},
	}
	*/

}
