package main

import "fmt"

func main() {
    numbers := []int{688, 99, 8, 109, 599}

    var min int = numbers[0]
    for _, value := range numbers {
        if value < min {
            min = value
        }
    }

    fmt.Println("En küçük sayı:", min)
}
