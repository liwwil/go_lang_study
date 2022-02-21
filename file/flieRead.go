package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.csv") //Open File
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")
		fmt.Println("Line ", count, ":", line)
		fmt.Println(item[2]) //Specific Item in CSV
		count++
	}
}
