package main

import "fmt"

var a int = b
var b int = f()
var c int = 2

func f() int {
	return c + 1
}

func main() {
	fmt.Print(a, b, c)
}
