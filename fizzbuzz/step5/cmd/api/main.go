package main

import (
	"log"
	"net/http"

	"github.com/nathancastelein/go-course-introduction/fizzbuzz/step5"
)

func main() {
	http.HandleFunc("/", step5.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
