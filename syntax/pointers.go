package syntax

import "fmt"

func Pointers() {
	myInt := 3
	myIntPointer := new(int)
	myIntPointer = &myInt

	fmt.Println("Pointer to int: ", myIntPointer)
	fmt.Println("Int from pointer: ", *myIntPointer)
}
