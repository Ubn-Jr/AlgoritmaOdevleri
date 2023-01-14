package main

import (
	"fmt"
	"sort"
)

func main() {

	var dDizi []int
	var bByt int = 5
	var dgr int
	fmt.Println("Dizi Boyutunu Giriniz")
	fmt.Scanln(&bByt)

	for i := 0; i < bByt; i++ {
		fmt.Printf("%d. Sayıyı Giriniz\n", i+1)
		fmt.Scanln(&dgr)
		dDizi = append(dDizi, dgr)
	}

	sort.Ints(dDizi)
	fmt.Printf("-----------------\n")
	for i := len(dDizi) - 1; i >= 0; i-- {
		fmt.Printf("%d\n", dDizi[i])
	}
}
