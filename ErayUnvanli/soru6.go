// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//BAŞARILI BİR UYGULAMA
func main() {

	var counts []int
	for i := 0; i < 3; i++ {
		fmt.Print("Sayı Giriniz : ")
		girilenSayi := 0
	    fmt.Scanf("%d", &girilenSayi)
		counts = append(counts, girilenSayi)
	}
	total:= 0
	for i := 0; i < len(counts); i++ {
		  total+=counts[i]
	}
	total=total/len(counts)
	fmt.Print(total)
}

