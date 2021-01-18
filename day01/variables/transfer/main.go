package main

import "fmt"

func main() {
	var1 := 10   // int
	var2 := 10.5 // float

	// 非法的
	// var3 := var1 + var2

	// 合法的
	var3 := var1 + int(var2) // var3 == 20

	fmt.Println(var3)
}
