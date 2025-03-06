package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(99)
	fmt.Println("My random number is:", rand.Intn(10))
}
