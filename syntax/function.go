package syntax

import "fmt"

func EmptyFunc() {
	fmt.Println("Hello from empty func")
}

func FuncWithInput(myString string) {
	fmt.Println("Hello from func with input", myString)
}

func FuncWithInputAndOuput(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

func FuncWithMultipleInputsAndOuputs(myString string, myInt1, myInt2 int) (string, int) {
	return fmt.Sprintf("Hello %s", myString), myInt1 + myInt2
}

func FuncWithNamedOutput() (result string) {
	result = "Hello there"
	return
}

func FuncSumWithVariadicInput(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func CallFuncWithVariadicInput() {
	result := FuncSumWithVariadicInput(1, 2, 3)
	fmt.Println("Result from sum function:", result)

	result = FuncSumWithVariadicInput(1, 2, 3, 4, 5)
	fmt.Println("Result from sum function:", result)

	numbers := []int{10, 20, 30}
	result = FuncSumWithVariadicInput(numbers...)
	fmt.Println("Result from sum function:", result)
}
