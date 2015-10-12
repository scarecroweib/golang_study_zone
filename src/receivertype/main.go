package main

import (
	"fmt"
)

const Pi float64 = 3.1415926

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type NameString string

func (r Rectangle) area() float64 { //ReceiverType r Rectangle是以值传递的,此处可以以指针传递，则会改变原值
	r.height = 100 //指针传递时会改变原值
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * Pi
}

func (s NameString) length() int { //ReciverType可以定义在任意自定义类型上 和 NameStringLength(s NameString)是等价的
	return len(s)
}

func NameStringLength(s NameString) int {
	return len(s)
}

func main() {
	r1 := Rectangle{12, 2}
	r2 := Rectangle{10, 10}
	c1 := Circle{10}
	c2 := Circle{12}

	fmt.Println(r1.area())
	fmt.Println(r2.area())
	fmt.Println(c1.area())
	fmt.Println(c2.area())

	fmt.Println(r1)
	fmt.Println(r2)

	var name NameString
	name = "weib"
	fmt.Println(name.length())

	fmt.Println(NameStringLength(name))
}
