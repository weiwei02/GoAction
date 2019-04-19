package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	const LENGTH = 10

	// 枚举
	const (
		a  = 1
		b2 = 2
	)

	// 枚举索引
	const (
		asn1 = iota
		asn2
		asn3
	)
}
