package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	// Echo command line arguments
	argsString := ""                // Short variable declaration
	//var argsString string         // Full variable declaration


	for i:=1; i<len(os.Args); i++ {
		argsString += os.Args[i] + " "
	}

	fmt.Println(argsString);

}