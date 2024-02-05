package step4

type FizzBuzz struct {
	number int
	result string
}

func NewFizzBuzz(number int) *FizzBuzz {
	return &FizzBuzz{
		number: number,
	}
}

func (f FizzBuzz) String() string {
	return f.result
}

// Write a Compute method for FizzBuzz.
// Compute calls fizzbuzz func on number field and store result in result field.
// Usage:
//     myFizzBuzz := NewFizzBuzz(5)
//     myFizzBuzz.Compute()
//     fmt.Println(myFizzBuzz.String()) // Should print "Buzz"
