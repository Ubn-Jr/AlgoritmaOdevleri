package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	var vize, final, ort float64
	var vizeMult float64 = 0.3
	var finalMult float64 = 0.7
	fmt.Println("Vize Notunuzu Giriniz")
	fmt.Scanln(&vize)

	fmt.Println("Final Notunuzu Giriniz")
	fmt.Scanln(&final)
	ort = (vize * vizeMult) + (final * finalMult)
	fmt.Printf("Ortalamanız %.1f\n", ort)

}
