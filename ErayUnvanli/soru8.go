package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print("Kilonuzu Giriniz : ")
	girilenSayiKilo := 0.0
	fmt.Scanf("%f", &girilenSayiKilo)

	fmt.Print("Boyunuzu Giriniz : ")
	girilenSayiBoy := 0.0
	fmt.Scanf("%f", &girilenSayiBoy)

	var kitkitleIndex float64
	kitleIndex := float64(girilenSayiKilo) / math.Pow(float64(girilenSayiBoy), 2)

	if kitleIndex < 18.5 {
		fmt.Println("Zayif")
	} else if kitleIndex >= 18.5 && kitkitleIndex <= 24.9 {
		fmt.Println("Normal")
	} else if kitleIndex >= 25 && kitkitleIndex <= 29.9 {
		fmt.Println("Fazla")
	} else if kitleIndex >= 30 && kitkitleIndex <= 34.9 {
		fmt.Println("1. derece obez")
	} else if kitleIndex >= 35 && kitkitleIndex <= 39.9 {
		fmt.Println("2. derece obez")
	} else {
		fmt.Println("Obez")
	}
	fmt.Println(kitleIndex)

	idealKilo := 25 * math.Pow(girilenSayiBoy, 2)

	fmt.Printf("Ideal kilonuz : %0.2v kg.\n", idealKilo)

}
