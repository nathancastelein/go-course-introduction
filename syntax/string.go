package syntax

import "fmt"

func String() {
	var myEmptyString string
	myString := "Hello"
	myStringWithDoubleQuotes := `"Nathan"`
	myMultiLineString := `Line 1
Line 2`

	fmt.Println("Here are some strings: ", myEmptyString, myString, myStringWithDoubleQuotes, myMultiLineString)
	fmt.Println("Size of a string: ", len(myString), len(myEmptyString))
	fmt.Println("Operation on a string: ", myString+" "+myStringWithDoubleQuotes)
}

func Fmt() {
	myAge := 33
	myName := "Nathan"
	mySize := 1.75

	myInformation := fmt.Sprintf("Hello, I'm %s, I am %d years old, and I am %.2fm tall", myName, myAge, mySize)
	fmt.Println(myInformation)
}
