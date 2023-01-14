package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------")
	var boy, kilo, endx float64
	var drm string
	fmt.Println("Kilonuzu Giriniz Kilogram")
	fmt.Scanln(&kilo)
	fmt.Println("Boyunuzu Giriniz Santimetre")
	fmt.Scanln(&boy)

	endx = kilo / ((boy / 100) * (boy / 100))
	switch {
	case endx < 18.5:
		drm = ("Zayıf")
	case endx >= 18.5 && endx < 24.9:
		drm = ("Normal")
	case endx >= 25 && endx < 29.9:
		drm = ("Fazla kilolu")
	case endx >= 30 && endx < 34.9:
		drm = ("I. derece obez")
	case endx >= 35 && endx < 39.9:
		drm = ("II. derece obez")
	case endx > 40:
		drm = ("III. derece obez")
	}

	fmt.Printf("Endeksiniz %.1f Durumunuz %s", endx, drm)

}
