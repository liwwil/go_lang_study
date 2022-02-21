package main //เมื่อวาง Defer ไว้หน้า Process ไหน Process นั้นจะทำงานท้ายสุด

import "fmt"

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Printf("The result of %f + %f is %f\n", value1, value2, result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println(j)
	}
}

func main() {
	/*fmt.Println("Calculator")
	defer fmt.Println("End")
	defer add(20, 10)
	defer add(15, 15)
	defer add(12, 12) // defer = Last it, firt out ตัวไหน defer บรรทัดสุดท้าย คือ ทำงานคนเเรก
	*/
	loop()
}
