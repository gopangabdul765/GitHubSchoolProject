package main

import "fmt"

func main() {
	x := 42
	y := x * 2 + 1
	z := y / 2
	if z > 0 {
		fmt.Println("The result is positive")
	} else if z < 0 {
		fmt.Println("The result is negative")
	} else {
		fmt.Println("The result is zero")
	}
}
