package main

import "fmt"

func main() {
	fmt.Println("=======soal 4========")
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	if int(d) > int(a) && int(d) > int(b) && int(d) > int(c) {
		fmt.Println("Ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}
}
