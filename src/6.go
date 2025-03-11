package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(42))

	// Generate a random number between 1 and 10, inclusive
	randomNumber := r.Intn(10) + 1

	fmt.Println("Random number:", randomNumber)
}
