package main

import "fmt"

func main() {
	// 十进制 类型推导出来默认int类型
	var a100 = 100
	// 十进制 int8 类型
	var a15 int8 = 15
	// 八进制
	var b024 = 024
	// 二进制
	var c5 = 0b101
	// 十六进制
	var d_5 = -0x5

	fmt.Printf("a100 => %v, Type => %T, transfer 0b:%b, o:%o, 0x:%x \n", a100, a100, a100, a100, a100)
	fmt.Printf("a15 => %v, Type => %T, transfer 0b:%b, o:%o, 0x:%x \n", a15, a15, a15, a15, a15)
	fmt.Printf("b024 => %v, Type => %T, transfer 0b:%b, o:%o, 0x:%x \n", b024, b024, b024, b024, b024)
	fmt.Printf("c5 => %v, Type => %T, transfer 0b:%b, o:%o, 0x:%x \n", c5, c5, c5, c5, c5)
	fmt.Printf("d_5 => %v, Type => %T, transfer 0b:%b, o:%o, 0x:%x \n", d_5, d_5, d_5, d_5, d_5)

	var b15 int8 = -15
	fmt.Printf("b15 => %v, transfer int32 : %d, 0x : %0x", b15, int32(b15), b15)
}
