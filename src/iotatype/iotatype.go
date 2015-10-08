package main

import "fmt"

func main() {
	const (
		x = iota
		y
		z

		x1 = iota
		y1
		z1
	)

	fmt.Println(z)
	fmt.Println(z1)
}
