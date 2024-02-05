package main

import (
	"log"
	"net/http"

	"github.com/nathancastelein/go-course-introduction/fizzbuzz/step6"
)

func main() {
	http.HandleFunc("/", step6.FizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
