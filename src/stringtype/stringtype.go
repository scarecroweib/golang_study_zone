package main

import "fmt"

func main() {
	var str1 = "hello world !"
	fmt.Println(str1)

	b1 := []byte(str1)
	b1[0] = 'H'
	str2 := string(b1)
	fmt.Println(str2)

	str3 := str1
	str3 = str3[:5] + " zhangjingwei"
	fmt.Println(str3)
}
