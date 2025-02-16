package step5

import (
	"fmt"
	"net/http"
	"strconv"
)

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	// Get query parameter "number" from URL
	numberString := r.URL.Query().Get("number")

	// Transform number from string to int
	number, err := strconv.Atoi(numberString)
	// If invalid input, write response header HTTP 400 then return
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Compute FizzBuzz on number
	result := fizzbuzz(number)

	// Write response header HTTP 200
	w.WriteHeader(http.StatusOK)
	// Write result on response
	fmt.Fprint(w, result)
}
