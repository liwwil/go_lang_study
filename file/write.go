package main

import "os"

func main() {
	data1 := []byte("hello\n Liu Lovely")
	err := os.WriteFile("data.txt", data1, 0644)

	if err != nil {
		panic(err)
	}

	f, err := os.Create("employeeName") //การสร้างไฟล์
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data2 := []byte("Liu\n Manee")
	os.WriteFile("employeeName", data2, 0644)
}
