package main

import (
	"fmt"
)

const (
	a = iota + 2024
	b
	c
	d
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v\t", a, b, c, d)
}
