package main

import "fmt"

func hitungJumlah(n int, jumlah *int) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya bilangan
	F.S jumlah yang menyatakan hasil penjumlahan digit kedua dan digit keempat dari n bilangan */
	var i, m int
	for i = 1; i <= n; i++ {
		fmt.Scan(&m)
		*jumlah = *jumlah + ((m % 10000) / 1000) + ((m % 100) / 10)

	}
}

func main() {
	var n, jumlah int
	fmt.Scan(&n)
	hitungJumlah(n, &jumlah)
	fmt.Println(jumlah)
}
