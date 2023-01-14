package main

import "fmt"

func main() {

	dDizi := []int{100, 100, 15, 100, 30, 20, 20}

	dSrlma := duplicate_count(dDizi)

	for kDgr, vTkrr := range dSrlma {
		fmt.Printf("%d Tane '%d' Elemanı, ", vTkrr, kDgr)
	}
	fmt.Printf("vardır.\n")

}
func duplicate_count(list []int) map[int]int {

	dup_freq := make(map[int]int) // iki boyutlu map int/int eğer kaç tane string olduğunu sayacak olsaydık string/int yapacaktık ve içeriye "(list []string) map[string]int" olarak alacaktık.

	for _, item := range list {
		// the dup_freq'ın icinde birden fazla kez var mı diye kontrol ediyoz
		_, exist := dup_freq[item]

		if exist {
			dup_freq[item] += 1 // varsa +1 arttırıyoruz
		} else {
			dup_freq[item] = 1 // yoksa 1'den başlatıyoruz.
		}
	}
	return dup_freq
}
