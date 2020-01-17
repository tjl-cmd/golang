package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "114.67.77.90:6379")
	if err != nil {
		fmt.Println("connect to redis error ", err)
		return
	}
	defer c.Close()
	if _, err := c.Do("AUTH", "123456"); err != nil {
		c.Close()
		fmt.Println(err)
		return
	}
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed1 ", err)
	}
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get 1:", err)
	} else {
		fmt.Printf("get mykey:%v\n", username)
	}
	//list, err := redis.String(c.Do("GET", "k1"))
	//if err != nil {
	//	fmt.Println("redis get failed2:", err)
	//} else {
	//	fmt.Printf("get mykey:%v\n", list)
	//}
}
