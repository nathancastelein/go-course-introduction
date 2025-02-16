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

func (f *FizzBuzz) Compute() {
	f.result = fizzbuzz(f.number)
}
