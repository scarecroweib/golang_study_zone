package main

import "fmt"

type Testint func(int) bool

func isUnder10(i int) bool {
	if i < 10 {
		return true
	}
	return false
}

func isBigThan10(i int) bool {
	if i >= 10 {
		return true
	}
	return false
}

func outputByFilter(f Testint, args ...int) {
	for _, v := range args {
		if f(v) {
			fmt.Println(v)
		}
	}
}

func main() {
	outputByFilter(isUnder10, 101, 1, 2, 3, 10, 100)
	outputByFilter(isBigThan10, 101, 1, 2, 3, 10, 100)
}
