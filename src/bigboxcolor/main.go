package main

import "fmt"

type Color byte

type Box struct {
	no                   string
	width, height, depth float64
	color                Color
}

type BoxList []Box

const (
	WHITH = iota
	BLACK
	BLUE
	RED
	YELLOW
	GREEN
)

//const ColorStrings = [6]string{"WHITH", "BLACK", "BLUE", "RED", "YELLOW", "GREEN"} 不能声明常量数组？？

func (b Box) Volume() float64 {
	return b.height * b.width * b.depth
}

//个人理解 在GO中，指针只需要在需要引用调用时写出来，其它情况编译器会自动处理，也就是只要告诉编译器是引用传递还是值传递就可以了
func (b *Box) SetColor(c Color) { //因为要改变原值，因此要使用指针
	//b.color = c //这里等价于下面一行
	(*b).color = c //使用指针时一定要加()
}

func (bl BoxList) BiggestBoxColor() Color {
	var v float64
	c := Color(WHITH) //将byte转换成Color类型
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			c = b.color
		}
	}
	return c
}

func (bl BoxList) PaintItBlack() {
	//	for _, b := range bl {
	//		b.SetColor(BLACK)
	//	}

	for i, _ := range bl {
		//(&bl[i]).SetColor(BLACK)	//两种写法等价
		bl[i].SetColor(GREEN)
	}
}

func (c Color) GetColorString() string {
	var colors = []string{"WHITH", "BLACK", "BLUE", "RED", "YELLOW", "GREEN"}
	return colors[c]
}

func main() {
	boxlist := BoxList{
		{"box1", 10, 10, 10, WHITH},
		{"box2", 20, 10, 10, BLACK},
		{"box3", 30, 10, 10, BLUE},
		{"box4", 40, 10, 10, RED},
		{"box5", 50, 10, 10, GREEN},
	}

	for _, box := range boxlist {
		fmt.Printf("%s---%f---%s\n", box.no, box.Volume(), box.color.GetColorString())
	}

	fmt.Printf("The biggest box's color: %s\n", boxlist.BiggestBoxColor().GetColorString())

	boxlist.PaintItBlack()

	for _, box := range boxlist {
		fmt.Printf("%s---%f---%s\n", box.no, box.Volume(), box.color.GetColorString())
	}
}
