package syntax

import "fmt"

func Slice() {
	var mySlice []int // len 0, cap 0
	printSlice(mySlice, "empty slice")

	myOtherSlice := []int{1, 2, 3} // len 3, cap 3
	printSlice(myOtherSlice, "non-empty slice")

	length := 5
	mySliceWithLength := make([]int, length) // len 5, cap 5
	printSlice(mySliceWithLength, "slice with length")

	capacity := 10
	mySliceWithMakeAndCapacity := make([]int, length, capacity) // len 5, cap 10
	printSlice(mySliceWithMakeAndCapacity, "slice with length and capacity")
}

func SliceOperations() {
	mySlice := []int{1, 2, 3}

	myElement := mySlice[0]                         // Accessing an element
	mySlice[2] = myElement * 3                      // Replace an element
	mySlice = append(mySlice, 4)                    // Append a new element at the end
	mySlice[0], mySlice[1] = mySlice[3], mySlice[2] // Swap multiple elements

	printSlice(mySlice, "slice after operations")
}

func SliceIteration() {
	mySlice := []int{1, 2, 3, 4}
	sliceSum := 0
	for i := 0; i < len(mySlice); i++ {
		sliceSum += mySlice[i]
	}
	fmt.Println("This is the sum of my slice: ", sliceSum)

	sliceMult := 1
	for i := range mySlice {
		sliceMult *= mySlice[i]
	}
	fmt.Println("This is the multiplication of my slice's elements: ", sliceMult)

	for i, element := range mySlice {
		fmt.Printf("This is the element %d from my slice: %d\n", i, element)
	}

	for _, element := range mySlice {
		fmt.Println("Iterating on my slice: ", element)
	}
}

func printSlice(slice []int, sliceName string) {
	fmt.Printf("Some information about %s: %d - Length %d - Capacity %d\n", sliceName, slice, len(slice), cap(slice))
}
