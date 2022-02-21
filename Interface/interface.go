package main //OOP

import (
	"fmt"
	"math"
)

type geometry interface { //Check ว่ามีฟังก์ชั่นอะไรบ้าง
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func Hello() {
	fmt.Println("Hello")
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area()) //ส่งค่าไปผ่าน Func area มันจะเเยกคำนวณ r and c เอง
	fmt.Println(g.perim())
}

func main() {
	Hello()
	r := rect{width: 3, height: 4} //ส่งค่าไปผ่าน Func measure
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
