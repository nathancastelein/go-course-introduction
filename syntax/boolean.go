package syntax

import "fmt"

func Boolean() {
	itsTrue, itsFalse := true, false

	fmt.Println("My booleans are: ", itsTrue, itsFalse)
	fmt.Println("Trying some operation: ", itsTrue && itsFalse, itsTrue || itsFalse)
}
