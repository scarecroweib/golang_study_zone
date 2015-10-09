package main

import (
	"fmt"
	funcs "functions"
)

func add1(a int32) int32 {
	return a + 1
}

func add1p(a *int32) {
	*a = *a + 1
}

func printArgs(arg ...int32) {
	fmt.Println("output start!!")
	for k, n := range arg {
		fmt.Printf("%d->%d\n", k, n)
	}
	fmt.Println("output end!")
}

func main() {
	var i32_1, i32_2 int32 = 1, 2
	result := funcs.Add(i32_1, i32_2)
	fmt.Println(result)

	sum1, pro1 := funcs.SumAndProduct(i32_1, i32_2)
	fmt.Printf("%d----%d\n", sum1, pro1)

	sum2, pro2 := funcs.SumAndProduct2(i32_1, i32_2)
	fmt.Printf("%d----%d\n", sum2, pro2)

	var a int32 = 100
	fmt.Printf("a:%d\n", a)

	fmt.Printf("a+1:%d, a:%d\n", add1(a), a)

	add1p(&a)
	fmt.Printf("*a+1 => a:%d\n", a)

	printArgs(1, 2, 3, 4, 5)

	//defer
	fmt.Println("defer start!")
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("defer run!")
}
