package main

import "fmt"

func init() {
	fmt.Println("declartion/main.go ==> init")
}

var myVersion = getVersion()

func main() {
	fmt.Println("declartion/main.go myVersion => " + myVersion)
}
