package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex  int
	Age  int
	//xxx  int
}

func main() {
	var s Student
	v := reflect.ValueOf(s)
	t := v.Type()
	//t := reflect.TypeOf(s)
	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Println("s is int64")
	case reflect.Struct:
		fmt.Println("s is struct")
		fmt.Printf("field num of s is %d\n", v.NumField())
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//fmt.Println(field)
			fmt.Printf("name:%s,type:%v,value%v\n", t.Field(i).Name, field.Type(), field.Interface())
		}
	default:
		fmt.Println("default")
	}

}
