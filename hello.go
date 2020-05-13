package main

import (
	"fmt"
)

func main()  {
	fmt.Println("hello!")
	x := 7

	if x > 6 {
		fmt.Println("More than 6")
	}

	// a := [5]int{5,4,3,2,1}  // Array
	// fmt.Println(a)

	a := []int{5,4,3,2,1} // Slice
	fmt.Println(a)

	a = append(a, 13)
	fmt.Println(a)

	// Map
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")
	fmt.Println(vertices)

	for i:= 0; i<5; i++ {
		fmt.Println(i)
	}

	for index, value := range a {
		fmt.Println("index: ", index, "value: ", value)
	}
}