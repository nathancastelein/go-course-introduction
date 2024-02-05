package step5

import "fmt"

var (
	fizz           = "Fizz"
	buzz           = "Buzz"
	fizzMultiplier = 3
	buzzMultiplier = 5
)

func fizzbuzz(input int) string {
	fizzBuzzValue := ""
	if isMultipleOf(input, fizzMultiplier) {
		fizzBuzzValue += fizz
	}

	if isMultipleOf(input, buzzMultiplier) {
		fizzBuzzValue += buzz
	}

	if fizzBuzzValue != "" {
		return fizzBuzzValue
	}

	return fmt.Sprintf("%d", input)
}

func isMultipleOf(number, divider int) bool {
	return number%divider == 0
}
