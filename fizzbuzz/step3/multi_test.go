package step3

import (
	"testing"
)

func TestMulti(t *testing.T) {
	// Arrange
	numbers := []int{1, 3, 5, 8, 15}
	expectedResults := []string{"1", "Fizz", "Buzz", "8", "FizzBuzz"}

	// Act
	results := multiFizzBuzz(numbers)

	// Assert
	if len(results) != len(expectedResults) {
		t.Fatalf("Length of result = %d; want %d", len(results), len(expectedResults))
	}

	for i := range expectedResults {
		if results[i] != expectedResults[i] {
			t.Errorf("Result element %d = %s; want %s", i, results[i], expectedResults[i])
		}
	}
}
