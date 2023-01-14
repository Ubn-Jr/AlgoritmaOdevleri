package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	var n, toplam int
	for n = 1; n <= 100; n++ {
		toplam += n
	}
	fmt.Printf("1'den 100 e Kadar sayıların Toplamı = %d 'dir\n", toplam)
}
