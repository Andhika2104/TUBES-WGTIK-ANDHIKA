package main

import "fmt"

func main() {
	fmt.Println("=======soal 3========")
	var x float64
	fmt.Scan(&x)

	if int(x)%3 == 0 {
		fmt.Println("Fizz")
	}
	if int(x)%5 == 0 {
		fmt.Println("Bazz")
	}
}
