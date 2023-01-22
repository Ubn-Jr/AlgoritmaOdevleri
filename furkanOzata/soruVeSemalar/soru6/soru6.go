package main

import "fmt"

func main() {
    numbers := []int{3, 7, 2, 9, 4}

    var total int = 0
    for _, value := range numbers {
        total += value
    }

    average := float64(total) / float64(len(numbers))

    fmt.Println("Ortalama:", average)
}