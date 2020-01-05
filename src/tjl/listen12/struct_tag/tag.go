package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string `json:"username"`
	Sex      string `json:"sex"`
	score    float32
}

func main() {
	user := &User{
		UserName: "user01",
		Sex:      "男",
		score:    99.2,
	}
	data, _ := json.Marshal(user)
	fmt.Printf("json str :%s\n", string(data))
	fmt.Printf("json str :%s\n", data)
}
