package main

import "fmt"

type Animal struct {
	Name string
	Sex  string
}

func (a *Animal) Talk() {
	fmt.Printf("i'talk i'm %s\n", a.Name)
}

type Dog struct {
	Feet string
	*Animal
}

func (d *Dog) Eat() {
	fmt.Println("dog is eat")
}
func (d *Dog) Talk() {
	fmt.Println("dog is talking1")
}

type puruAnimal struct {
}

func (p *puruAnimal) Talk() {
	fmt.Println("buru dongwu talk")
}
func main() {
	var d *Dog = &Dog{
		Feet: "four feet",
		Animal: &Animal{
			Name: "dog",
			Sex:  "xiong",
		},
	}
	d.Name = "dog"
	d.Sex = "xiong"
	d.Eat()
	d.Talk()
}
