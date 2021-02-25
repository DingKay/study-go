package main

import "fmt"

const (
	a = iota // a == 0
	b = iota // b == 1
	c = iota // c == 2
	d        // d == 3 ( 相当于 d = iota )
)
const (
	a1 = iota + 1 // iota =0; a == 0 + 1
	_             // iota =1;(implicitly _ = iota + 1 )
	b1            // iota =2;b == 3 (implicitly b = iota + 1 )
	c1            // iota =3;c == 4 (implicitly c = iota + 1 )
)

const (
	t1 = iota;t2 = iota
	t3 = iota;t4 = iota
)

func main() {
	fmt.Println(a1, b1, c1)
	fmt.Println(t1,t2,t3,t4)
}
