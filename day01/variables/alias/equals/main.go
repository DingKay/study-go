package main

import "fmt"

type float float64

func main() {
	var f float = 52.2
	var g float64 = 52.2

	fmt.Println("f == g", f == g)
}
