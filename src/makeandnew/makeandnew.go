package main

import "fmt"

func main() {
	var fslice []int32
	fslice = make([]int32, 10)
	fslice = append(fslice, 100)
	fmt.Println(fslice)

	var testmap map[string]int32
	testmap = make(map[string]int32)
	fmt.Println(testmap)

	testmap2 := make(map[string]int32)
	testmap2["test1"] = 1
	testmap2["test2"] = 2
	fmt.Println(testmap2)

	//new
}
