package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var arr = [...]int{1, 3, 5, 7, 8}
	var count int
	for _, v := range arr {
		count += v
	}
	fmt.Println(count)
	findSumIndex(arr, 8)

	// var testArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// slice1 := testArr[2:6:9]
	// fmt.Printf("%v, len=> %d, cap=> %d \n", slice1, len(slice1), cap(slice1))

	// 1, 2, 3, 4, 5, 6, 7, 8, 9
	// 1, 4, 5, 6, 7, 8, 9, 8, 9
	var testArr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := append(testArr[0:1], testArr[3:]...)
	fmt.Printf("slice1=>%v, len=> %d, cap=> %d \n", slice1, len(slice1), cap(slice1))
	fmt.Printf("testArr=>%v, len=> %d, cap=> %d \n", testArr, len(testArr), cap(testArr))

	args := [...]int{1, 3, 5}

	// a := args[0:1:1]
	// b := args[2:3:3]
	// fmt.Printf("a = %v, len = %d, cap = %d\n", a, len(a), cap(a))
	// fmt.Printf("b = %v, len = %d, cap = %d\n", b, len(b), cap(b))
	// res := append(a, b...)

	res := append(args[0:1], args[2:3:3]...)
	fmt.Printf("res = %v, len = %d, cap = %d\n", res, len(res), cap(res))

	var inta = new(int)
	fmt.Printf("%+v \n", *inta)

	var arrInt = new([]int)
	fmt.Printf("%+v\n", *arrInt)

	var varNewMap = new(map[string]string)
	fmt.Printf("%+v,type=>%T\n", varNewMap, varNewMap)

	var varMakeMap = make(map[string]string)
	fmt.Printf("%+v,type=>%T\n", varMakeMap, varMakeMap)

	n1 := make([]int, 10)
	fmt.Println("result", n1)
	fmt.Println(len(n1), cap(n1))
	fmt.Println(n1[2])

	fmt.Println("====================")

	n2 := new([]int)
	fmt.Println("result", n2, *n2)
	fmt.Println(len(*n2), cap(*n2))
	// fmt.Println((*n2)[2]) panic

	var resultMap = make(map[string]int)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		resultMap[fmt.Sprintf("T%02d", i)] = rand.Intn(100)
	}
	fmt.Println(resultMap)
	var sortKey = make([]string, 10, 10)
	var index = 0
	for k := range resultMap {
		// sortKey = append(sortKey, k)
		sortKey[index] = k
		index++
	}
	sort.Strings(sortKey)
	fmt.Println(sortKey)
	for i := 0; i < len(sortKey); i++ {
		fmt.Print(sortKey[i], ":", resultMap[sortKey[i]], "; ")
	}

}

func findSumIndex(arr [5]int, sum int) {
	for i, v := range arr {
		for j := i + 1; j < len(arr); j++ {
			if count := v + arr[j]; count == sum {
				fmt.Printf("%d,%d\n", i, j)
			}
		}
	}
}
