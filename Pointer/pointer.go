package main

import "fmt"

func zeroVal(ivalue int) {
	ivalue = 0
}

func zeroPointer(ipointer *int) { // กำหนดประเภท Pointer เป็น int
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i", i)
	zeroVal(i) //Func ไม่ใช้ Pointer i ไม่เปลี่ยนค่า
	fmt.Println("i from func zero : ", i)

	zeroPointer(&i) // & คือเข้าถึงที่อยู่ของตัวแปร
	fmt.Println("i from func zeroPointer : ", i)
	fmt.Println("i address from func zeroPointer : ", &i)

}
