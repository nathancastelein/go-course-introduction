package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello there wonderful people!")
	fmt.Printf("Here is your id: %s\n", uuid.New())
}
