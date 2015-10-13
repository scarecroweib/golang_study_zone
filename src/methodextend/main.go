package main

import (
	f "fmt"
)

type Human struct {
	name string
	age  int
}

func (h Human) sayHi() {
	f.Println("Hello, I am a Human!")
}

type Student struct {
	Human
	grade int
}

func (s Student) sayHi() {
	f.Println("Hello, I am a Student!")
}

func main() {
	s := Student{Human{"weib", 12}, 6}

	f.Println(s)

	s.sayHi()

	s2 := new(Student)
	s2.name = "weib2"
	(*s2).age = 100 //使用指针时一定要加()，即(*s2)
	f.Println(*s2)
}
