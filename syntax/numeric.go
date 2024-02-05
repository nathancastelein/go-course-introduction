package syntax

import "fmt"

func Numeric() {
	var myInt64 int64
	myInt := 10
	myFloat64 := 3.14
	myInt64 = 20

	fmt.Println("Here are numeric types: ", myInt, myInt64, myFloat64)
	fmt.Println("With some operation: ", myInt+int(myInt64), float64(myInt)+myFloat64, myInt-int(myFloat64))
}
