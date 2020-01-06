package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name" db:"db"`
	Sex   int     `json:"sex"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
	//xxx  int
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	t := v.Type()

	filed := t.Elem().Field(0)
	fmt.Printf("tag json=%s\n", filed.Tag.Get("json"))
	fmt.Printf("tag json=%s\n", filed.Tag.Get("db"))

}
