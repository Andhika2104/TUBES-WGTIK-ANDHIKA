package main

import "fmt"

func main() {
	fmt.Println("========soal 5=========")
	var teks string
	fmt.Scan(&teks)

	for teks != "selesai" {
		fmt.Println(teks)
		fmt.Scan(&teks)
	}
}
