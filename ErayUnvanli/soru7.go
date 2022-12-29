// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//BAŞARILI BİR UYGULAMA
func main() {
	fmt.Print("Vize Notu Giriniz : ")
	midtermExam := 0
	fmt.Scanf("%d", &midtermExam)
	fmt.Print("Final Notu Giriniz : ")
	finalExam := 0
	fmt.Scanf("%d", &finalExam)
	var midtermAndFinalCount int
	midtermAndFinalCount = control(midtermExam, finalExam)
	fmt.Println(midtermAndFinalCount)
}

func control(midtermExam int, finalExam int) int {
		midterm := (midtermExam * 30) / 100
		final := (finalExam * 70) / 100
		return midterm + final
}

