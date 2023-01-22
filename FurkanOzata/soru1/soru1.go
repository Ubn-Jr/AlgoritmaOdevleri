package main

import "fmt"

func main() {
	n := 70
	m := 30
	total := control(n, m)
	fmt.Println(total)
}


func control(n int, m int) int {
return n + m
}