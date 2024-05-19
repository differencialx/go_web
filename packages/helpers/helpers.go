package helpers

import "math/rand"

func RandomNumber(n int) int {
	return rand.Intn(n)
}
