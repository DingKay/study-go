package main

import "fmt"

const constant = "constant value"

const (
	constant1 = "constant value1"
	constant2 = "constant value2"
)

const (
	a = 1 // a==1
	b = 2 // b==2
	c     // c==2
	d     // d==2
)

func main() {
	fmt.Println(a, b, c, d)
}
