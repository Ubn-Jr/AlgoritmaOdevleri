package main

import (
	"fmt"
)

func main() {

	var counts []int
	for i := 0; i < 7; i++ {
		fmt.Print("Sayı Giriniz : ")
		girilenSayi := 0
		fmt.Scanf("%d", &girilenSayi)
		counts = append(counts, girilenSayi)
	}

	for i := 0; i < len(counts); i++ {
		deger := false
		for l := 0; l < i; l++ {
			if counts[l] == counts[i] {
				deger = true
			}
		}
		if deger == false {
			count := 0
			for c := 0; c < len(counts); c++ {
				if counts[c] == counts[i] {
					count++
				}
			}
			fmt.Println(counts[i], "sayisindan", count, "adet var")
		}
	}
}
