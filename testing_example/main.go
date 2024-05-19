package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(100, 10)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result:", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Cannot devide by 0")
	}

	result = x / y
	return result, nil
}
