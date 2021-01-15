package main

import (
	"fmt"

	"github.com/dingkay/study/day01/declaration/version"
)

func init() {
	fmt.Println("declartion/entry.go ==> init")
}

func getVersion() string {
	fmt.Println("declartion/entry.go getVersion")
	return version.Version
}
