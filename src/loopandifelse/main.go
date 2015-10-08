package main

import "fmt"

func main() {
	loopIndex := 10
	for i := 0; i <= loopIndex; i++ {
		fmt.Printf("Loop:%d \n", i)
	}

	fmt.Println("=============")

	var index int = 0
	for index <= loopIndex {
		fmt.Printf("while:%d \n", index)
		index++

		if index > 5 {
			goto jump
		}
	}

jump:
	fmt.Println("######loop end!")

	var testmap map[string]int32
	testmap = make(map[string]int32)
	testmap["test1"] = 1
	testmap["test2"] = 2

	xslice := []int32{1, 2, 3}

	for k, v := range testmap {
		fmt.Printf("%s=>%d\t", k, v)
	}
	fmt.Println()

	for _, v := range testmap {
		fmt.Println("map's value:", v)
	}

	for _, v := range xslice {
		fmt.Println(v)
	}

	goto switch_

switch_:
	selected := 100
	switch selected {
	case 1:
		fmt.Println("switch 1")
	case 2:
		fmt.Println("switch 2")
	case 3:
		fmt.Println("switch 3")
	case 100:
		fmt.Println("switch 100")
	default:
		fmt.Println("switch x")
	}
}
