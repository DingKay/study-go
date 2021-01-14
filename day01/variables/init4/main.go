package main

import "fmt"

var a int = b
var b int = c
var c int = a

func main() {
	fmt.Print(a, b, c)
}
