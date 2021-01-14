package main

import (
	"fmt"

	"github.com/dingkay/study/day01/pkgtest/pkg1"
	"github.com/dingkay/study/day01/pkgtest/pkg1/pkg2"
)

func main() {
	fmt.Println(pkg1.PkgMsg)
	fmt.Println(pkg2.PkgMsg)
}
