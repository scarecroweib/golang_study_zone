package main

import (
	"fmt"
)

const G_PI float64 = 3.1415926

func main() {
	const Pi float32 = 3.1415926

	var i32test int32
	i32test = 32
	fmt.Println(i32test)

	var i32test1, i32test2, i32test3 int32
	i32test1 = 1
	i32test2 = 2
	i32test3 = 3
	var i32test4_1 = 199
	fmt.Printf("%d, %d, %d, %d\n", i32test1, i32test2, i32test3, i32test4_1)

	i32test4 := 4
	i32test5, i32test6 := 5, 6
	fmt.Printf("%d, %d, %d\n", i32test4, i32test5, i32test6)

	var r float32 = 10.0
	fmt.Println(Pi * r * r)

	var r2 float64 = 20.0
	fmt.Println(G_PI * r2 * 2)

	var i32test1_1, i32test1_2, i32test1_3 int32 = 1, 2, 3
	fmt.Printf("%d, %d, %d\n", i32test1_1, i32test1_2, i32test1_3)

	var btest1 bool
	if !btest1 {
		btest1 = true
	}
	fmt.Println(btest1)
}
