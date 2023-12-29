package main

import (
	"fmt"
)

func main() {
	fmt.Println("Composite types")

	// Arrays have fixed size
	var array [3]int = [3]int{10, 15, 20}
	fmt.Printf("%d %d %d \n", array[0], array[1], array[2])
	fmt.Printf("array[2]=%d \n", array[len(array)-1])

	for index, value := range array {
		fmt.Printf("array[%d]=[%d] \n", index, value)
	}

	type MyEnum int

	const (
		VAL_1 = iota
		VAL_2
	)

	// [...] = syntax that uses array in size of number of init values (in this case 2)
	enumStrings := [...]string{VAL_1: "VAL_1", VAL_2: "VAL_2"}
	fmt.Println(VAL_1, enumStrings[VAL_1])

	// Slices
	var sl []int
	for i := 0; i < 7; i++ {
		// Append will increase the slice size.
		// The size will be multiplied by 2:
		// 0: cap=[0]
		// 1: cap=[1]
		// 2: cap=[2]
		// 3: cap=[4]
		// 4: cap=[4]
		// 5: cap=[8]
		// 6: cap=[8]

		fmt.Printf("%d: cap=[%d] \n", i, cap(sl))
		sl = append(sl, i)
	}

	// Maps
	ages := make(map[string]int)
	ages["ahmed"] = 56
	ages["ara"] = 62

	for key, value := range ages {
		fmt.Println(key, value)
	}

	delete(ages, "ahmed")

	age, ok := ages["ahmed"]
	if !ok {
		fmt.Printf("ahmed in not a key in ages map \n")
	} else {
		fmt.Printf("ahmed's age is %d \n", age)
	}

	// Structs
	type Person struct {
		ID   int
		Name string
	}

	var person1 Person
	person1.ID = 12
	person1.Name = "Ahmed"

	personPtr := &person1
	fmt.Printf("Person ID=[%d], name=[%s] \n", personPtr.ID, personPtr.Name)
}
