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
	enk:= counts[0]
	for i := 0; i < len(counts); i++ {
		  if (counts[i] < enk){
			enk = counts[i];
		  }
	}
	fmt.Print(enk)
}

