package step5

import (
	"fmt"
	"net/http"
)

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameter "number" from URL

	// Transform number from string to int
	// If invalid input, write response header HTTP 400 then return

	// Compute FizzBuzz on number

	// Write response header HTTP 200
	// Write result on response
	fmt.Fprintf(w, "Hello World")
}
