package syntax

import "fmt"

func Declare() {
	var myFirstBoolean bool
	mySecondBoolean := true
	myThirdBoolean := myFirstBoolean
	myFourstBoolean := new(bool)

	fmt.Println("My variables are: ", myFirstBoolean, mySecondBoolean, myThirdBoolean, myFourstBoolean)
}

func ZeroValue() {
	var myBoolean bool
	var myInt int
	var myString string

	fmt.Println("Zero values are: ", myBoolean, myInt, myString)
	// print "Zero values are:  false 0 "
}
