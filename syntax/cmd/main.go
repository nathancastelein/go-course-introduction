package main

import (
	"fmt"

	"github.com/nathancastelein/go-course-introduction/syntax"
)

func main() {
	syntax.Declare()
	syntax.ZeroValue()
	syntax.Boolean()
	syntax.Numeric()
	syntax.String()
	syntax.Fmt()
	syntax.Array()
	syntax.Slice()
	syntax.SliceOperations()
	syntax.SliceIteration()
	syntax.Map()
	syntax.MapOperations()
	syntax.MapIteration()
	syntax.EmptyFunc()
	syntax.FuncWithInput("Nathan")
	fmt.Println("Output from func: ", syntax.FuncWithInputAndOuput("Nathan"))
	myString, myInt := syntax.FuncWithMultipleInputsAndOuputs("Nathan", 1, 2)
	fmt.Println("Multiple outputs: ", myString, myInt)
	syntax.CallFuncWithVariadicInput()
	syntax.Pointers()
	syntax.If()
	syntax.For()
	syntax.Switch()
	syntax.InitializeStruct()
}
