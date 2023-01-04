package main

import "fmt"

func main() {
	fmt.Println("=======soal 1========")
	var S string
	var hasil_penjumlahan, A, B int
	fmt.Scanln(&S, &A, &B)
	hasil_penjumlahan = A + B
	fmt.Println("kata =", S)
	fmt.Println("Jumlah =", hasil_penjumlahan)
}
