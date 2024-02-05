package step4

import "testing"

func TestFizzBuzzStruct(t *testing.T) {
	// Arrange
	myFizzBuzz := NewFizzBuzz(5)

	// Act
	myFizzBuzz.Compute()

	// Assert
	if myFizzBuzz.String() != "Buzz" {
		t.Errorf("NewFizzBuzz(5).Compute = %s; want Buzz", myFizzBuzz.String())
	}
}
