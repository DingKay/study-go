package main

import "fmt"

var arr [10]int

func init() {
	fmt.Println("main init arr")
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
}

func main() {
	fmt.Println("main.go ==> start")
	fmt.Println(arr)
}
