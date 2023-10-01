package main

import (
	"fmt"
)

var x int = 1998
var y string = "clinton"
var z bool = true

func main() {
	s := fmt.Sprintf("%d %s %t", x , y, z)
	fmt.Println(s)
}
