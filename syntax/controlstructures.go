package syntax

import "fmt"

func If() {
	if 2 > 1 {
		fmt.Println("Two is superior to one")
	} else if 1+1 == 11 {
		fmt.Println("1 + 1 equals 11")
	} else {
		fmt.Println("We're in the else statement")
	}

	if sum := FuncSumWithVariadicInput(1, 2, 3); sum > 5 {
		fmt.Println("The computed sum was superior to 5: ", sum)
	}
}

func For() {
	for myVariable := 0; myVariable < 5; myVariable++ {
		fmt.Println("I'm in a loop: ", myVariable)
	}

	for myVariable := range 5 {
		fmt.Println("I'm in a loop: ", myVariable)
	}

	for range 5 {
		fmt.Println("I'm in a loop")
	}

	myVariable := 0
	for myVariable == 0 {
		myVariable += 5
	}

	for {
		myVariable++
		if myVariable > 10 {
			fmt.Println("Stopping the loop, value: ", myVariable)
			break
		}
	}
}

func Switch() {
	name := "Nathan"
	switch name {
	case "Nicolas":
		fmt.Println("Hello Nicolas!")
	case "Nathalie":
		fmt.Println("Hello Nathalie!")
	case "Nathan":
		fmt.Println("Hello Nathan!")
	default:
		fmt.Println("Who are you? Who who, who who?")
	}

	myInt := 10
	switch {
	case myInt > 20:
		fmt.Println("What an int!")
	case myInt+5 > 15:
		fmt.Println("Let's add some compute")
	case myInt >= 10:
		fmt.Println("My int is not so big: ", myInt)
	}
}
