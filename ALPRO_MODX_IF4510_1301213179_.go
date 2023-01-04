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

	fmt.Println("========soal 2=========")
	var r int
	var luas_lingkaran float64

	fmt.Scan(&r)
	luas_lingkaran = 22.0 / 7.0 * float64(r) * float64(r)
	fmt.Println("Luas lingkaran (r =", r, ") adalah", luas_lingkaran)

	fmt.Println("========soal 3=========")
	var x float64
	fmt.Scan(&x)

	if int(x)%3 == 0 {
		fmt.Println("Fizz")
	}
	if int(x)%5 == 0 {
		fmt.Println("Bazz")
	}

	fmt.Println("=======soal 4========")
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	if int(d) > int(a) && int(d) > int(b) && int(d) > int(c) {
		fmt.Println("Ada rekor baru")
	} else {
		fmt.Println("Tidak ada rekor baru")
	}

	fmt.Println("========soal 5=========")
	var teks string
	fmt.Scan(&teks)

	for teks != "selesai" {
		fmt.Println(teks)
		fmt.Scan(&teks)
	}

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

	fmt.Println("========soal 7==========")
	var a, b, c, d, e, f int
	var x int
	fmt.Scanln(&a, &b, &c)
	fmt.Scanln(&d, &e, &f)
	fmt.Scanln(&x)

	if x == 0 || x == 360 {
		fmt.Println(a, b, c)
		fmt.Println(d, e, f)
	} else if x == 90 {
		fmt.Println(d, a)
		fmt.Println(e, b)
		fmt.Println(f, c)
	} else if x == 180 {
		fmt.Println(f, e, d)
		fmt.Println(c, b, a)
	} else if x == 270 {
		fmt.Println(c, f)
		fmt.Println(b, e)
		fmt.Println(a, d)
	}

	fmt.Println("=========soal 8=============")
	var r, t int
	fmt.Scan(&r, &t)
	fmt.Println(hitungVolume(r, t))

	fmt.Println("==========soal 9============")
	var n int
	var t1, t2, t3 int
	fmt.Scanln(&n)
	jumlah = 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&t1, &t2, &t3)
		if (t1 + t2 + t3) >= 2 {
			jumlah = jumlah + 1
		}
	}
	fmt.Println(jumlah)
}
