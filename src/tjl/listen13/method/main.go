package main

import "fmt"

type People struct {
	Name    string
	Country string
}

func (p People) Print() {
	fmt.Printf("name=%s country:%s\n", p.Name, p.Country)
}
func (p People) Set(name, country string) {
	p.Name = name
	p.Country = country
}
func (p *People) SetV2(name, country string) {
	p.Country = country
	p.Name = name
}
func main() {
	var p1 People = People{
		Name:    "people01",
		Country: "china",
	}
	p1.Print()
	p1.Set("people02", "enligsh")
	(&p1).SetV2("people02", "enligsh")
	p1.SetV2("people03", "enligsh")
	p1.Print()
}
