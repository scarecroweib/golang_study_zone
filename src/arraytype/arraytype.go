package main

import "fmt"

func main() {
	//array
	var ia_1 [10]int32
	ia_1[0] = 100
	ia_1[1] = 101
	ia_1[2] = 102
	ia_1[3] = 103
	ia_1[4] = 104
	fmt.Println(ia_1[3])

	var ia_2 = [...]int32{1, 2, 3, 4, 5}
	fmt.Println(ia_2[2])

	var ia_3 = [2][2]int32{[2]int32{1, 2}, {4, 5}}
	fmt.Println(ia_3[1][1])

	//slice
	var fslice []int32 //只声明一个指针不初始化
	//fslice = {1,2,3} 不能再次赋值（理解为声明完成后，内存不能再次申请）
	//fslice[0] = 1
	//fslice[1] = 2
	//fslice[2] = 3
	fslice = ia_1[0:4] //指向一个已存在的数组，但是并没有复制，只是引用，所以当数据变化时，元数据也发生变化，因为他们本来就是一个
	fslice[2] = 200
	fmt.Println(fslice)
	fmt.Println(ia_1)

	xslice := []int{1, 2, 3} //初始化一个slice
	fmt.Println(xslice)

	yslice := append(xslice, 100)
	fmt.Println(yslice)
	fmt.Println(xslice)
	fmt.Printf("yslice len:%d\n", len(yslice))
	fmt.Printf("yslice cap:%d\n", cap(yslice))

	wslice := append(fslice, 1000, 1001, 1002, 1003, 1004, 1005, 1006, 1007) //当一次追加超过原数组上限时，新声明空间，不再修改原有数组
	fmt.Println(wslice)
	fmt.Println(fslice)
	fmt.Println(ia_1)
	fmt.Printf("wslice len:%d\n", len(wslice))
	fmt.Printf("wslice cap:%d\n", cap(wslice))

	fmt.Println("##########################################")

	zslice := append(append(append(append(append(append(append(fslice, 300), 400), 500), 600), 700), 800), 900)
	fmt.Println(zslice)
	fmt.Println(fslice)
	fmt.Println(ia_1)
	fmt.Printf("zslice len:%d\n", len(zslice))
	fmt.Printf("zslice cap:%d\n", cap(zslice))

	//map(非线程安全)
	var xmap map[string]int
	xmap = make(map[string]int)
	xmap["name"] = 100
	fmt.Println(xmap["name"])

	ymap := make(map[string]int32)
	ymap["name"] = 200
	fmt.Println(ymap)
}
