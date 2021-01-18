package main

import "fmt"

var var1, var2, var3 int = 3, 4, 5
var var4, var5, var6 = 1, 2.2, false

func main() {
	// 标准声明
	var bol bool = true
	// 类型推导（根据值判断该变量是什么类型）
	var str = "string type"
	// 简短声明
	a := "a word"
	// 批量声明
	var (
		test1        = "test1"
		test2 string = "test2"
	)
	var7, var8, var9 := 1, 2, 3
	fmt.Println(bol)
	fmt.Println(str)
	fmt.Println(a)
	fmt.Println(test1 + test2)
	fmt.Println(var7 + var8 + var9)
}
