package main

import "fmt"

func hitungVolume(r, t int) float64 {
	pi := 3.14
	volume := float64(r) * float64(r) * pi * float64(t)
	return volume
}

func main() {
	fmt.Println("=========soal 8=============")
	var r, t int
	fmt.Scan(&r, &t)
	fmt.Println(hitungVolume(r, t))
}
