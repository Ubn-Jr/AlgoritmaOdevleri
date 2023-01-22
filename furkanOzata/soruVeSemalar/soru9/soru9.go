package main

import (
	"fmt"
	"sort"
)

func main() {
    arr := []int{3, 6, 1, 9, 2, 5}
    sort.Sort(sort.IntSlice(arr))
    fmt.Println(arr)
}