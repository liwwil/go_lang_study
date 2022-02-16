package main

import "fmt"

var productName [4]string
var price [4]float32 //ประกาศ Array ชนิด Float

func main() {
	productName[0] = "Iphone" // assign ค่าให้กับ เเต่ละ Index ของ ProductName Array
	productName[1] = "Ipad"
	productName[2] = "Macbook"
	productName[3] = "Macbook Pro"

	price2 := [4]float32{1000, 2000, 3000, 4000} //ประกาศ Array ชนิด Float แบบเต็ม
	fmt.Println(productName)
	fmt.Println(price)
	fmt.Println(price2)
}
