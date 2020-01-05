package main

import "fmt"

type Animal interface {
	Talk()
	Eat()
	Name() string
}
type Dog struct {
}

func (d *Dog) Talk() {
	fmt.Println("汪汪汪")
}
func (d *Dog) Eat() {
	fmt.Println("我在啃骨头")
}
func (d *Dog) Name() string {
	str := "我叫joke"
	return str
}

type Pig struct {
}

func (p *Pig) Talk() {
	fmt.Println("啃啃啃")
}
func (p *Pig) Eat() {
	fmt.Println("我在吃草")
}
func (p *Pig) Name() string {
	str := "我叫猪八戒"
	return str
}
func main() {
	var d Dog
	var a Animal
	a = &d
	a.Eat()
	a.Talk()
	name := a.Name()
	fmt.Println(name)
	var pig Pig
	a = &pig
	fmt.Println(a.Name())
	a.Talk()
	a.Eat()
}
