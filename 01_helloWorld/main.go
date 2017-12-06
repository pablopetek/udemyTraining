package main

import "fmt"

// var declaration
var b string

const p string = "asdf"

func main() {

	const q = 42
	const (
		ac = 333.14
		aa = "asdf"
	)

	const (
		_  = iota             // 0
		KB = 1 << (iota * 10) // 1 << (1 * 10)
		MB = 1 << (iota * 10) // 1 << (2 * 10)
	)

	// shorthand declaration
	// and assignment of variables (initialization)

	a := 10
	b := "golang"
	c := 4.17
	d := true
	fmt.Printf("%v \t %T \n", a, a)
	fmt.Printf("%v \t %T \n", b, b)
	fmt.Printf("%v \t %T \n", c, c)
	fmt.Printf("%v \t %T \n", d, d)

	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)

}
