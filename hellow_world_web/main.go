package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// Home home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page hit")
	fmt.Fprintf(w, "This is the home page")
}

// Abount is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	log.Println("About page hit")
	fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// addValues adds two values
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 0.0)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y

	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
