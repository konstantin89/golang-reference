package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

type ColoredPoint struct {
	// Note that here we store Point value.
	// We can store Point* to allow methods that receive Point* from Point struct.
	Point
	Color int
}

// Print is method of Point struct
// (p Point) is called RECEIVER. Its the instance that the method is performed on.
// Be aware that here p is copy of Point!
func (p Point) Print() {
	fmt.Printf("(%d, %d) \n", p.X, p.Y)
}

// Pointer RECEIVER allows modifications in the given instances
func (p *Point) ScaleBy(factor int) {
	p.X *= factor
	p.Y *= factor
}

// moveByOne is private method to this package
// Method/Field visibility is controller by the first letter (capital or not)
func (p *Point) moveByOne() {
	p.X++
	p.Y++
}

// Interface example
type Printed interface {
	Print()
}

func main() {
	p := Point{1, 3}
	p.Print()

	p.ScaleBy(3)
	p.Print()

	p.moveByOne()
	p.Print()

	c := ColoredPoint{Point{2, 4}, 255}
	// Compiler automatically generates wrapper method for ColoredPoint that receives ColoredPoint and calls Point.Print().
	c.Print()

	// Using an interface
	var printed Printed = &p
	printed.Print()

}
