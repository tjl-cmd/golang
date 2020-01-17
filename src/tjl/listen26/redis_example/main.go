package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "114.67.77.90:6379")
	if err != nil {
		fmt.Println("conn redis failed", err)
		return
	}
	defer c.Close()
	if _, err := c.Do("AUTH", "123456"); err != nil {
		c.Close()
		fmt.Println(err)
		return
	}
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed", err)
		return
	}
	fmt.Println(r)

}
