package syntax

import "fmt"

func Array() {
	var myIntArray [3]int
	myIntArray[0] = 3
	myIntArray[1] = 2
	myIntArray[2] = 1

	fmt.Println("Here is an int array", myIntArray)
}
