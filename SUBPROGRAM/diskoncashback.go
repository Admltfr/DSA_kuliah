package main

import "fmt"

func diskon(totalbelanja int) int {

	/*Mengembalikan angka 1 jika belanja
	pembeli memenuhi syarat diskon dan
	angka 0 jika tidak memenuhi syarat*/
	var jumlah int
	if (totalbelanja/100)%4 == 0 {
		jumlah = 1
	} else {
		jumlah = 0
	}
	return jumlah

}

func cashback(totalbelanja int) int {

	/*Mengembalikan angka 1 jika belanja pembeli
	memenuhi syarat cashback dan
	angka 0 jika tidak memenuhi syarat*/
	var jumlah int
	if (totalbelanja/100)%3 == 0 {
		jumlah = 1
	} else {
		jumlah = 0
	}
	return jumlah
}

func main() {
	var banyaknyaPembeli, totalBelanja, i int
	var DapatDiskon, DapatCashback int
	fmt.Scan(&banyaknyaPembeli)
	DapatDiskon = 0
	DapatCashback = 0
	for i = 1; i <= banyaknyaPembeli; i++ {
		fmt.Scan(&totalBelanja)
		DapatDiskon = DapatDiskon + diskon(totalBelanja)
		DapatCashback = DapatCashback + cashback(totalBelanja)
	}
	fmt.Println(DapatDiskon, DapatCashback)
}
