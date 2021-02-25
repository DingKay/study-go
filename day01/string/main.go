package main

import (
	"fmt"
	"strings"
)

func main() {
	var t1 string = "中文"
	var t2 string = "en"
	fmt.Println(t1, t2)

	var s1 string = "c:\\User\\go"
	fmt.Println(s1)

	var s2 string = `
	c:\User\go
	hello
	`
	fmt.Println(s2)

	var stringlen = "12345"
	fmt.Println(len(stringlen))

	// 字符串的拼接
	var name = "kay"
	var job = "kick"

	s3 := name + job
	fmt.Println(s3)

	s4 := fmt.Sprintf("%s%s", name, job)
	fmt.Println(s4)

	fmt.Printf("%s%s \n", name, job)

	// 分割
	splitStr := strings.Split(s1, "\\")
	fmt.Printf("%T\t%v\n", splitStr, splitStr)

	// join 拼接
	fmt.Println(strings.Join(splitStr, "+"))
}
