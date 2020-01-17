package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

//初始化连接池
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		Dial: func() (r redis.Conn, err error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { //测试连接是否可用
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         64,                //最小连接数
		MaxActive:       1000,              //最大活跃连接数
		IdleTimeout:     240 * time.Second, //连接超时
		Wait:            false,
		MaxConnLifetime: 0,
	}
}
func main() {
	pool = newPool("114.67.77.90:6379", "123456")
	for {
		time.Sleep(time.Second)
		conn := pool.Get()
		conn.Do("set", "abc", 10000)

		r, err := redis.Int(conn.Do("get", "abc"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(r)
	}
}
