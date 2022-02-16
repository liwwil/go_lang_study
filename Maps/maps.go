//Maps จะเกี่ยวข้องกับ Key and Value

package main

import "fmt"

var product = make(map[string]float64) //***map[key type]value type

func main() {
	fmt.Println("product:", product)

	//add product
	product["Iphone"] = 1000
	product["Ipad"] = 20000
	product["Ipod"] = 2000
	fmt.Println("product:", product)
	//delete product
	delete(product, "Ipod")
	fmt.Println("product:", product)
	//update product
	product["Iphone"] = 40000
	fmt.Println("product:", product)

	// เข้าถึง Value ของ Maps
	value1 := product["Iphone"]
	fmt.Println("value1:", value1)

	//กำหนดค่าเริ่มต้นให้ตัวแปร Maps
	courseName := map[string]string{"101": "Golang", "102": "Java", "103": "Python"}
	fmt.Println("courseName:", courseName)
}
