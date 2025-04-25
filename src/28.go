package main

import (
    "fmt"
)

func main() {
    // Example of using pointers to manipulate arrays in Go
    arr := []int{1, 2, 3, 4, 5}
    fmt.Println(arr[10])  // Output: 2 (Index out of range)
}

// This function demonstrates how to use pointers in Go to access elements within a slice.
