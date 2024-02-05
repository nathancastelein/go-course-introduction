package step3

import "testing"

func TestMultiMap(t *testing.T) {
	// Arrange
	numbers := []int{1, 3, 5, 8, 15}
	expectedResults := map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		8:  "8",
		15: "FizzBuzz",
	}

	// Act
	results := multiFizzBuzzMap(numbers)

	// Assert
	if len(results) != len(expectedResults) {
		t.Fatalf("Length of result = %d; want 6", len(results))
	}

	for expectedKey, expectedValue := range expectedResults {
		value := results[expectedKey]
		if expectedValue != value {
			t.Errorf("Result element %d = %s; want %s", expectedKey, value, expectedValue)
		}
	}
}
