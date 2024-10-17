package main

import "fmt"

func uangKakak(x int, n int) int {
	/*
		I.S. terdefinisi dua buah bilangan bulat yang menyatakan tabungan di minggu pertama dan lamanya menabung (dalam minggu).
		F.S. total tabungan kakak setelah n minggu.
	*/

	var i, jumlah, minggu int
	jumlah = 0
	minggu = n
	for i = 1; i <= n; i++ {
		minggu--
		jumlah = jumlah + x + (minggu * 2500)
	}
	return jumlah
}

func main() {
	var tabPertama, n, jumlah int
	fmt.Scan(&tabPertama, &n)
	jumlah = uangKakak(tabPertama, n)
	fmt.Println(jumlah)
}
