package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "114.67.77.90:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	if _, err := c.Do("AUTH", "123456"); err != nil {
		c.Close()
		fmt.Println(err)
		return
	}
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 2000)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("rpop", "book_list"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
