package step5

import "testing"

func TestShouldReturnSameNumberWhenItIsNotMultipleOfThreeOrFive(t *testing.T) {
	// Arrange
	number := 1

	// Act
	result := fizzbuzz(number)

	// Assert
	if result != "1" {
		t.Errorf("FizzBuzz(1) = %s; want 1", result)
	}
}

func TestShouldReturnFizzWhenNumberIsMultipleOfThree(t *testing.T) {
	// Arrange
	number := 3

	// Act
	result := fizzbuzz(number)

	// Assert
	if result != "Fizz" {
		t.Errorf("FizzBuzz(3) = %s; want Fizz", result)
	}
}

func TestShouldReturnBuzzWhenNumberIsMultipleOfFive(t *testing.T) {
	// Arrange
	number := 5

	// Act
	result := fizzbuzz(number)

	// Assert
	if result != "Buzz" {
		t.Errorf("FizzBuzz(5) = %s; want Buzz", result)
	}
}

func TestShouldReturnFizzBuzzWhenNumberIsMultipleOfThreeAndFive(t *testing.T) {
	// Arrange
	number := 15

	// Act
	result := fizzbuzz(number)

	// Assert
	if result != "FizzBuzz" {
		t.Errorf("FizzBuzz(15) = %s; want FizzBuzz", result)
	}
}
