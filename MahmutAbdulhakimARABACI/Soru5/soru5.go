package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	n := [6]int{13, 11, 7, 5, 3, 2}
	var i, ksayi int
	ksayi = n[0]
	for i = 0; i < len(n); i++ {
		if ksayi > n[i] {
			ksayi = n[i]
		}
	}
	fmt.Printf("Dizideki En Kücük Sayi %d'dir\n", ksayi)
}
