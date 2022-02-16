//Func มี 4 ประเภท
//1. เป็น Function ไม่รับค่าเเละ Return ค่า
//2. เป็น Function มีการรับค่า
//3. เป็น Function มีการ Return ค่า
//4. เป็น Function มีการรับค่าและ Return ค่า

package main

import "fmt"

func hello() {
	fmt.Println("Hello") //ไม่รับค่า เเละ ไม่Return ค่าออกไป
}

func plus(value1 int, value2 int) int { //รับค่า เเละ Return ค่ากลับ เเละก็อย่าลืมกำหนดประเภทของค่าที่ Return
	return value1 + value2
}

func plusThree(val1, val2, val3 int) int {
	return val1 + val2 + val3
}

func main() {
	//เรียกใช้งาน Function
	hello()
	result := plus(1, 2)
	fmt.Println("result:", result)
	result2 := plusThree(1, 2, 3)
	fmt.Println("result2:", result2)
}
