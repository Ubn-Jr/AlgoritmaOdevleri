package main

import "fmt"

func main() {
    var vize,final float64

    fmt.Print("Vize Notunu Giriniz : ")
    fmt.Scan(&vize)

    fmt.Print("Final Notunu Giriniz : ")
    fmt.Scan(&final)

    vize = vize * 0.3
    final = final * 0.7
    ortalama := vize + final

    fmt.Println("Not Ortalaması : ",ortalama)
}