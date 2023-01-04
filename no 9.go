package main

import "fmt"

func main() {
	fmt.Println("==========soal 9============")
	var n int
	var t1, t2, t3 int
	fmt.Scanln(&n)
	jumlah := 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&t1, &t2, &t3)
		if (t1 + t2 + t3) >= 2 {
			jumlah = jumlah + 1
		}
	}
	fmt.Println(jumlah)
}
