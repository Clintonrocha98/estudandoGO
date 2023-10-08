package main

import (
	"fmt"
)

const (
	a = 1
	b = "testando"
	c = 2.34
)
const x int = 1
const y string = "tstando"
const z float64 = 2.34

func main() {
	fmt.Printf("%v - %T\n%v - %T\n%v - %T\n", a, a, b, b, c, c)
	fmt.Printf("%v - %T\n%v - %T\n%v - %T\n", x, x, y, y, z, z)

}
