package main

import "fmt"

func main() {

	courseName := []string{"Golang", "Nodejs", "Python", "Java"}
	fmt.Println(courseName)
	courseName = append(courseName, "C++", "Ruby", "Java")
	fmt.Println(courseName)

	courseWeb := courseName[:4] //ประกาศตัวแปร
	fmt.Println(courseWeb)

	courseWeb = courseName[2:4] //ไม่ต้องประกาศซ้ำ Assign ค่าเข้าได้เลย
	fmt.Println(courseWeb)
}
