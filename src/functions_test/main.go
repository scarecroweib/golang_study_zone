package main

import (
	"fmt"
	funcs "functions"
)

func main() {
	var i32_1, i32_2 int32 = 1, 2
	result := funcs.Add(i32_1, i32_2)
	fmt.Println(result)
}
