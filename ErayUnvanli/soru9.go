package main

import (
	"fmt"
	"sort"
)

func main() {

	var counts []int
	for i := 0; i < 3; i++ {
		fmt.Print("Sayı Giriniz : ")
		girilenSayi := 0
		fmt.Scanf("%d", &girilenSayi)
		counts = append(counts, girilenSayi)
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})
	for _, v := range counts {
		fmt.Println(v)
	}
}
