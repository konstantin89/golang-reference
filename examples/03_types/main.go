package main

import (
	"fmt"
	"math"
)

func main() { 

	// +++ Integer s+++
	var u uint8 = 255
	var i int8 = -1

	//s := u + i     // This fails compilation (invalid operation: u + i (mismatched types uint8 and int8))
	s := int(u) + int(i)

	fmt.Printf("u=[%d], i=[%d], u+i=[%d] \n", u, i, s);

	// +++ Floating points +++
	var pi float32 = math.Pi
	fmt.Printf("Max float32 = [%f], f=[%f] \n", math.MaxFloat32, pi);

	// +++ Strings (are immutable) +++
	str := "string"
	for _, c := range str {
		fmt.Printf("%c\n", c)      // Access by range
	}

	for i:=0; i<len(str); i++{
		fmt.Printf("%c", str[i])   // Access by index
	}
	fmt.Printf("\n")

	// +++ Constants +++
	const const_var = 16

	fmt.Printf("c_var=[%d] \n", const_var);

	type Weekday int // This is a way to define enums (with iota generator)
	const (
		Sunday Weekday = iota // iota generator (starts from 0 and increases)
		Monday
		Tuesday
		//...
	)

	fmt.Printf("Sunday=[%d], Monday=[%d], Tuesday=[%d] \n", Sunday, Monday, Tuesday);
	// stdout:  Sunday=[0], Monday=[1], Tuesday=[2]
}
