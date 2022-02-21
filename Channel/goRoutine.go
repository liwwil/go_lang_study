package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	go f("Hello world") //Goroutine = Process ย่อยๆ + ทำงานเร็วมากจนไม่ Print เราต้องใส่ Time Delay
	go f("message")
	time.Sleep(5 * time.Second)

}
