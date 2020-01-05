package main

import "fmt"

type Employer interface {
	CalcSalary() float32
}
type Programer struct {
	name  string
	base  float32
	extra float32
}

func NewProgramer(name string, base float32, extra float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extra: extra,
	}
}
func (p *Programer) CalcSalary() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extra float32
}

func NewSale(name string, base float32, extra float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extra: extra,
	}
}
func (s *Sale) CalcSalary() float32 {
	return s.base + s.base*0.5
}
func calcAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost += v.CalcSalary()
	}
	return cost
}
func main() {
	p1 := NewProgramer("搬砖", 10000, 0)
	p2 := NewProgramer("搬砖1", 10000, 0)
	p3 := NewProgramer("搬砖2", 10000, 0)

	s1 := NewSale("销售1", 2000, 4000)
	s2 := NewSale("销售2", 2000, 4000)
	s3 := NewSale("销售3", 2000, 4000)
	var employList []Employer
	employList = append(employList, &p1)
	employList = append(employList, &p2)
	employList = append(employList, &p3)
	employList = append(employList, &s1)
	employList = append(employList, &s2)
	employList = append(employList, &s3)
	cost := calcAll(employList)
	fmt.Printf("这个月人力总成本为:%f", cost)
}
