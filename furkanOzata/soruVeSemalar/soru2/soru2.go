package main

import "fmt"

func main() {
	fmt.Print("Sayı Gir : ")
	enteredNumber := 0
	fmt.Scanf("%d", &enteredNumber)
	if enteredNumber%2 == 0{
		fmt.Println("Sayi ciftir")
	}	else {
		fmt.Println("Sayi tektir")
	}
	
}