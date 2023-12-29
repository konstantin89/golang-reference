package main

import (
	"fmt"
)

func hello_world() {
	fmt.Println("Hello World")
}

// valName is optional. Mostly for documentation.
func multiple_returns() (valName1 int, valName2 int) {
	return 3, 4
}

func add(x int, y int) (sum int) {
	s := x + y
	return s
}

func main() {
	hello_world()
	fmt.Printf("%T \n", hello_world)

	var x, y int
	x, y = multiple_returns()
	fmt.Printf("x=%d, y=%d \n", x, y)

	func_value := hello_world
	func_value()

	add(3, 2)
}
