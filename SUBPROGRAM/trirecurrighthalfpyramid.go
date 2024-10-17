package main

import "fmt"

func segitigarecursive2(baris int, kolom int, n int) {
	/*  I.S Terdefinisi nilai bilangan bulat baris, kolom, dan n
	F.S menampilkan pola  string "*" yang berbentuk segitiga */
	if kolom <= n {
		if kolom <= n-baris { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan,

			fmt.Print(" ")
		} else {
			fmt.Print("*")
		}
		segitigarecursive2(baris, kolom+1, n) // gunakan fungsi rekursif pada baris ini

	} else if baris < n {
		fmt.Println()
		segitigarecursive2(baris+1, 1, n) // gunakan fungsi rekursif pada baris ini
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	segitigarecursive2(1, 1, n)
}
