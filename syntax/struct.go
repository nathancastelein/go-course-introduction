package syntax

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func InitializeStruct() {
	jessie := Person{
		Name: "Jessie",
		Age:  17,
	}

	james := Person{"James", 17}

	fmt.Println("Jessie's age is ", jessie.Age)

	fmt.Println("Famous Team Rocket members: ", jessie, james)
}

func (p *Person) HappyBirthday() {
	p.Age = p.Age + 1
}

func (p Person) YearOfBirth() int {
	return time.Now().Year() - p.Age
}

type UniquePerson struct {
	cardIdNumber string // Is not visible outside the package
	Name         string // Visible and modifiable outside the package
	Age          int    // Visible and modifiable outside the package
}

func NewUniquePerson(cardIdNumber, name string, age int) *UniquePerson {
	return &UniquePerson{
		cardIdNumber: cardIdNumber,
		Name:         name,
		Age:          age,
	}
}
