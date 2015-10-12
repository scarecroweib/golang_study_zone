package main

import "fmt"

type Color byte

type Box struct {
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

func (b *Box) SetColor(c Color) { //因为要改变原值，因此要使用指针
	b.color = c
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
		bl[i].SetColor(BLACK)
	}
}

func (c Color) GetColorString() string {
	var colors = []string{"WHITH", "BLACK", "BLUE", "RED", "YELLOW", "GREEN"}
	return colors[c]
}

func main() {
	boxlist := BoxList{
		{10, 10, 10, WHITH},
		{20, 10, 10, BLACK},
		{30, 10, 10, BLUE},
		{40, 10, 10, RED},
		{50, 10, 10, GREEN},
	}

	for _, box := range boxlist {
		fmt.Printf("%f---%s\n", box.Volume(), box.color.GetColorString())
	}

	fmt.Printf("The biggest box's color: %s\n", boxlist.BiggestBoxColor().GetColorString())

	boxlist.PaintItBlack()

	for _, box := range boxlist {
		fmt.Printf("%f---%s\n", box.Volume(), box.color.GetColorString())
	}
}
