package main

import "fmt"

type float float64

func main() {
	var f float = 52.2

	fmt.Printf("f has value %v and type %T", f, f)
}
