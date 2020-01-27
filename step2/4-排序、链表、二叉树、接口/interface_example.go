package main

import (
	"fmt"
	"math"
)

// 打印不同几何图形的面积和周长

type geometry interface {
	area() float32
	perim() float32
}
type rect struct {
	length, width float32
}
type circle struct {
	radius float32
}
func (r *rect) area() float32{
	return  r.length * r.width
}
func (r *rect) perim() float32{
	return 2 * (r.length + r.width)
}
func (c *circle) area() float32{
	return math.Pi * c.radius * c.radius
}
func (c *circle) perim() float32{
	return 2 * math.Pi * c.radius
}
func show(name string, param interface{}){
	switch param.(type){
	case geometry:
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type!")
	}
}

func main() {
	rec := &rect{
		length: 1,
		width: 2,
	}
	cir := &circle{
		radius: 1,
	}
	//var g geometry
	//g = rec
	//fmt.Println(g.area())
	show("rect", rec)
	show("circle", cir)
	show("test", "test param")
}