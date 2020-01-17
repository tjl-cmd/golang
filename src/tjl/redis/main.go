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
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed ", err)
	}
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("get mykey:%v\n", username)
	}
	list, err := redis.String(c.Do("GET", "k1"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("get mykey:%v\n", list)
	}
}
