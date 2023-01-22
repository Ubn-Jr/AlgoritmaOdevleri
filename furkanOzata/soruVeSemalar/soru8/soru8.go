package main

import (
	"fmt"
	"math"
)

func main() {
    var kilo, boy float64

    fmt.Print("Kilonuzu kg olarak giriniz : ")
    fmt.Scan(&kilo)

    fmt.Print("Boyunuzu m olarak giriniz : ")
    fmt.Scan(&boy)

    vki := kilo / (boy * boy)

    if vki < 18.5 {
        fmt.Println("VKI : ", vki , "Zayıf")
    } else if vki < 24.9 {
        fmt.Println("VKI : ", vki , "Normal")
    } else if vki < 29.9 {
        fmt.Println("VKI : ", vki , "Fazla kilolu")
    } else if vki < 34.9 {
        fmt.Println("VKI : ", vki , "I. derece obez")
    } else if vki < 39.9 {
        fmt.Println("VKI : ", vki , "II. derece obez")
    } else {
        fmt.Println("VKI : ", vki , "III. derece obez")
    }

	idealKilo := 25 * math.Pow(boy, 2)

	fmt.Printf("Ideal kilonuz : %0.2v kg.\n", idealKilo)
}