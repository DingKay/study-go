package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("A number from 1-100", rand.Intn(81))
	for i := 0; i < 100; i++ {
		fmt.Println(rand.Intn(10000))
	}
}
