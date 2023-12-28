package main

import (
	"fmt"
)

func main() { 
	fmt.Println("Composite types");
	
	// Arrays have fixed size
	var array[3]int =[3]int{10,15,20}
	fmt.Printf("%d %d %d \n", array[0], array[1], array[2]);
	fmt.Printf("array[2]=%d \n", array[len(array)-1]);

	for index,value := range array{
		fmt.Printf("array[%d]=[%d] \n", index, value);
	}

	type MyEnum int

	const (
		VAL_1 = iota
		VAL_2
	)

	// [...] = syntax that uses array in size of number of init values (in this case 2)
	enumStrings := [...]string{VAL_1:"VAL_1", VAL_2:"VAL_2"}
	fmt.Println(VAL_1, enumStrings[VAL_1]);

	// Slices

}