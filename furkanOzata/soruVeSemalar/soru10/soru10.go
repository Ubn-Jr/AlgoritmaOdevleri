package main

import (
	"fmt"
)

func main() {
    arr := []int{3, 3, 2, 3, 4, 1}
    counts := make(map[int]int)

    for _, v := range arr {
        counts[v]++
    }

    for key, value := range counts {
        fmt.Printf("%d tane %d elemanı vardır\n", value, key)
    }
}