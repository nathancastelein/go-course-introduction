package main

import (
	"log"
	"net/http"

	"github.com/nathancastelein/go-course-introduction/fizzbuzz/step4"
)

func main() {
	http.HandleFunc("/", step4.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
