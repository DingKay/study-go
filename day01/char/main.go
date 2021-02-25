package main

import "fmt"

func main() {
	c1 := '中'
	c2 := 'z'
	fmt.Printf("%T \t %v \t %x \n", c1, c1, c1)
	fmt.Printf("%T \t %v \t %x \n", c2, c2, c2)
	traversalString()
}

func traversalString() {
	str := "你好，我是golang"

	// byte --> uint8
	fmt.Printf("%T \t %v \n", str[0], str[0])
	for i := 0; i < len(str); i++ {
		// fmt.Printf("%v(%c) ", str[i], str[i])
		fmt.Printf("%X ", str[i])
	}
	fmt.Println()
	s := []rune(str)[0]
	// rune --> int32
	fmt.Printf("%T \t %v \n", s, s)
	for _, char := range str {
		fmt.Printf("%v(%c) ", char, char)
	}
}
