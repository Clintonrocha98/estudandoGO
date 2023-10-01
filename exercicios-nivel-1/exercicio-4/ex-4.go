package main

import (
	"fmt"
)

type teste int

var x teste

func main() {
	fmt.Println("initial value: ", x)
	fmt.Printf("inital type: %T \n", x)
	x = 42
	fmt.Println("New value: ",x)

}
