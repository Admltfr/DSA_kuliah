package main

import "fmt"

type dataPeserta [100]int

func main() {
	var arrPeserta dataPeserta
	var pertama, kedua int
	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ { //memasukkan data
		fmt.Scan(&arrPeserta[i])
	}
	pertama = arrPeserta[0]
	for i = 1; i < n; i++ { // cek nilai terbesar dan terbesar kedua
		if pertama < arrPeserta[i] {
			kedua = pertama
			pertama = arrPeserta[i]
		} else if kedua < arrPeserta[i] {
			kedua = arrPeserta[i]

		}
	}
	fmt.Println("Kedua terbesar bernilai :", kedua)
}
