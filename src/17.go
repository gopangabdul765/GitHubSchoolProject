package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    numbers := make([]int, 10)
    for i := range numbers {
        numbers[i] = rand.Intn(10) + 1
    }

    fmt.Println(numbers)
}
