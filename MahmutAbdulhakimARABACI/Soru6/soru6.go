package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	n := [6]int{13, 11, 7, 5, 3, 2}
	var i, toplam int
	var ort float32

	for i = 0; i < len(n); i++ {
		toplam += n[i]
	}
	ort = float32(toplam) / float32(len(n))
	fmt.Printf("Dizideki Sayıların Ortalaması %.1f'dir\n", ort)
}
