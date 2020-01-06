package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Sex   int
	Age   int
	Score float32
	//xxx  int
}

func (s *Student) SetName(name string) {
	s.Name = name
}
func (s *Student) Print() {
	fmt.Printf("通过反射进行调用%v\n", s)
}
func main() {
	var s Student
	s.SetName("xxxx")
	v := reflect.ValueOf(&s)
	t := v.Type()
	fmt.Printf("struct student have %d methods\n", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		//fmt.Println(method)
		fmt.Printf("struct %d method,name:%s type:%v\b\n", i, method.Name, method.Type)
	}
	//通过reflect.value获取对应的方法并调用
	m1 := v.MethodByName("Print")
	var args []reflect.Value
	m1.Call(args)

	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value
	name := "stu01"
	namev := reflect.ValueOf(name)
	args2 = append(args2, namev)
	m2.Call(args2)
	m1.Call(args)
}
