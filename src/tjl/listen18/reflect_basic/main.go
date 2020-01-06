package main

import (
	"fmt"
	"reflect"
)

func reflect_example(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fmt.Printf("type of a is %v\n", v)
	fmt.Printf("type of a is :%v\n", t)
	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64")
	case reflect.String:
		fmt.Printf("a is string")
	}
}
func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	//t := reflect.TypeOf()一致
	//t := v.Type()
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64,store value is:%d", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float,store value is:%f", v.Float())
	case reflect.String:
		fmt.Printf("a is string,store value is:%s", v.String())
	case reflect.Complex64:
		fmt.Printf("a is Complex64,store value is:%v", v.Complex())
	}
}
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		v.SetInt(89)
		fmt.Printf("a is int64,store value is:%d", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float,store value is:%f", v.Float())
	case reflect.String:
		fmt.Printf("a is string,store value is:%s", v.String())
	case reflect.Complex64:
		fmt.Printf("a is Complex64,store value is:%v", v.Complex())
	case reflect.Ptr:
		fmt.Println("set a to 89")
		v.Elem().SetInt(89)
	default:
		fmt.Println("default")
	}
}
func main() {
	var x int64 = 23
	//reflect_example(x)
	//reflect_value(x)
	reflect_set_value(&x)
	fmt.Println(x)
}
