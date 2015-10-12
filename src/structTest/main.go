package main

import (
	f "fmt"
)

type skills []string

type person struct {
	name string //非声明变量，所以不加var
	age  int32
}

type monkey struct {
	name string
	age  int32
}

type student struct {
	person //嵌套
	monkey //当嵌套的结构体，包含相同的字段时
	skills
	gread int32
}

func Older(p1, p2 person) (name string, agediff int32) {
	if p1.age > p2.age {
		name = p1.name
		agediff = p1.age - p2.age
	} else {
		name = p2.name
		agediff = p2.age - p1.age
	}
	return
}

func main() {
	var s1 *person = new(person) //指针声明
	s1.age = 20
	s1.name = "weib"

	f.Println(s1)
	f.Println(*s1)

	var s2 person //声明
	s2.age = 21
	s2.name = "XX"
	f.Println(s2)

	s3 := person{"xx1", 22}            //声明并初始化（按顺序）
	s4 := person{age: 23, name: "xx2"} //声明并初始化（指定字段）
	f.Println(s3)
	f.Println(s4)

	var oldername string
	var agediff int32

	oldername, agediff = Older(s2, s3)
	f.Printf("s2 and s3:%s(%d)\n", oldername, agediff)

	oldername, agediff = Older(s2, s4)
	f.Printf("s2 and s4:%s(%d)\n", oldername, agediff)

	oldername, agediff = Older(s3, s4)
	f.Printf("s3 and s4:%s(%d)\n", oldername, agediff)

	var ss1 student
	ss1.monkey.name = "weib" //当嵌套的结构体，包含相同的字段时，如要操作则必须指明
	ss1.gread = 6
	ss1.monkey.age = 12

	//skills作为一个slice存在，在此操作前都没有申请内存或初始化
	//以下是slice的两种初始化方法
	//ss1.skills = []string{"skill1", "skill2"}
	ss1.skills = make([]string, 10)
	ss1.skills[0] = "skill1"
	ss1.skills = append(ss1.skills, "skill3")
	f.Printf("skills len:%d\n", len(ss1.skills))
	f.Println(ss1)

	ss2 := student{person{"tom", 13}, monkey{"monkey king", 12}, []string{"skill2"}, 4} //嵌套形式下，声明并初始化
	f.Println(ss2)

}
