package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Soru 2 Tek mi Çift mi?------")
	fmt.Println("---------------------------------")
	var n int
	fmt.Println("Sayıyı Giriniz")
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Printf("Sayı %d Çift Bir Sayıdır.\n", n)
	} else {
		fmt.Printf("Sayı %d Tek Bir Sayıdır.\n", n)
	}
}
