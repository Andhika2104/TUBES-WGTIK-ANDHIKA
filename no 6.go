package main

import "fmt"

func main() {
	fmt.Println("========soal 6=========")
	var nilai int
	jumlah := 0
	n := 0
	fmt.Scanln(&nilai)
	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scanln(&nilai)
	}
	var rata2 float64
	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah) / float64(n)
	}
	fmt.Print(rata2)
}
