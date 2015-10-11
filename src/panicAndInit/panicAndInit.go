package main

import (
	. "fmt" //.的定义是可以在此包内调用这个包的方法时，不再显示指定包名
)

var index int32 //包级变量

const max = 10

//包中的初始化方法会自动执行
func init() {
	Println("go init")
	index = 5
}

func testIndex() {
	Println("go testIndex")
	for ; index < 100; index++ {
		if index < max {
			Println(index)
		} else {
			panic("index is over max") //相当于抛出异常，在GO中是进入恐慌
		}
	}
}

func main() {
	x := recover() //进程没有在恐慌状态时，返回nil
	Println(x)

	defer func() { //defer的定义要在可能进入恐慌的代码之前，要不然不会执行
		if y := recover(); y != nil { //recover()相当于捕获异常，在GO中是将进程从恐慌状态恢复
			Println("panic has be catched!")
			Println(y) //会输出进入恐慌时的信息 "index is over max"
		}
	}()

	testIndex()
}
