// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//BAŞARILI BİR UYGULAMA
func main() {
	fmt.Print("Sayı Giriniz : ")
	girilenSayi := 0
	fmt.Scanf("%d", &girilenSayi)
	if girilenSayi%2 == 0{
		fmt.Println("Sayi cift")
	}	else {
		fmt.Println("Sayi tek")
	}
	
}

