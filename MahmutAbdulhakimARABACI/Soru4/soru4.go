package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	var n int
	for n = 1; n <= 100; n++ {
		if n%2 == 1 {
			fmt.Printf("Tek Sayı: %d\n", n)
		}
	}
}
