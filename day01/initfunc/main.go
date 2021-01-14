package main

import (
	"fmt"
)

func init() {
	fmt.Println("main.go ==> init() [1]")
}

func init() {
	fmt.Println("main.go ==> init() [2]")
}

func init() {
	fmt.Println("main.go ==> init() [3]")
}

func main() {
	fmt.Println("main.go ==> main")
}
