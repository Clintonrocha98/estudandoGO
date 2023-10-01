package main

import (
	"fmt"
)

type teste int

var x teste
var y int

func main() {
	fmt.Println("initial value: ", x)
	fmt.Printf("inital type: %T \n", x)
	x = 42
	fmt.Println("New value: ",x)
	y = int(x)
	fmt.Println("y value: ", y)
	fmt.Printf("y type: %T \n", y)
}
