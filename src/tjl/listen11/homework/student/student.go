package main

import (
	"fmt"
	"math/rand"
)

func testInterface() {
	var a interface{}
	var b int = 100
	var c float32 = 1.3
	var d string = "hello"

	a = b
	fmt.Printf("a=%v\n", a)
	a = c
	fmt.Printf("a=%v\n", a)
	a = d
	fmt.Printf("a=%v\n", a)
}
func studentStore() {
	var stuMpa map[int]map[string]interface{}
	stuMpa = make(map[int]map[string]interface{}, 16)
	//插入学生id=1，姓名=stu01，分数=78.2，年龄= 18
	var id = 1
	var name = "stu01"
	var score = 78.2
	var age = 18
	value, ok := stuMpa[id]
	if !ok {
		value = make(map[string]interface{}, 8)
	}
	value["name"] = name
	value["id"] = id
	value["score"] = score
	value["age"] = age
	stuMpa[id] = value
	for i := 0; i < 10; i++ {
		value, ok := stuMpa[i]
		if !ok {
			value = make(map[string]interface{}, 8)
		}
		value["name"] = fmt.Sprintf("sut%d", i)
		value["id"] = i
		value["score"] = rand.Float32() * 100
		value["age"] = rand.Intn(100)
		stuMpa[i] = value
	}
	for key, v := range stuMpa {
		fmt.Printf("id:%d stu info = %#v\n", key, v)
	}
	//fmt.Printf("stuMap:%#v\n", stuMpa)
}
func main() {
	//testInterface()
	studentStore()
}
