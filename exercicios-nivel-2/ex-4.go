package main

import (
	"fmt"
)

func main() {
	a := 31337
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\t\n", a, a, a)
	fmt.Println("------------------------------")
	fmt.Printf("%d\t%b\t%#x\t", b, b, b)

}
