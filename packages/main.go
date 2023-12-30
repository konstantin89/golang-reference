package main

// go mod init my-package
// go get github.com/guptarohit/asciigraph
// go test ./tests

import (
	"fmt"
	// Importing package that is child directory
	// In go, there is max one package per directory
	"my-package/hello"

	"github.com/guptarohit/asciigraph"
)

func DrawExampleGraph() {
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
}

func main() {
	fmt.Println("main is starting!")

	hello.HelloPackage()

	DrawExampleGraph()
}
