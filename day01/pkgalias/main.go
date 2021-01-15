package main

import (
	"fmt"

	_ "github.com/dingkay/study/day01/pkgalias/greet"
	subgreet "github.com/dingkay/study/day01/pkgalias/greet/greet"
)

func main() {
	fmt.Println(subgreet.Message)
}
