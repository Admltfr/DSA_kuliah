package main

import "fmt"

const NMax int = 100

type tabChar [NMax]byte

func masukanArray(T *tabChar, n *int) {
	/*I.S. Data tersedia pada piranti masukan
	  Proses: Masukan akan terus berlangsung dan berhenti ketika pengguna
	  memasukkan '.'
	  F.S. Sekumpulan karakter sebanyak n berada dalam array T
	  Petunjuk: Gunakan while-loop untuk melakukan proses input
	*/
	var input byte
	*n = 0
	fmt.Scanf("%c", &input)
	for input != '.' {
		(T)[*n] = input
		*n++
		fmt.Scanf("%c", &input)
	}
}

func lowerCase(T *tabChar, n int) {
	/*I.S. Terdefinisi sekumpulan n karakter pada array T
	  F.S. Seluruh anggota karakter pada array T dikonversi menjadi huruf kecil
	  Petunjuk: Gunakan ASCII, apabila sudah huruf kecil, tidak perlu diubah
	*/
	var i int
	for i = 0; i < n; i++ {
		if T[i] >= 65 && T[i] <= 90 {
			T[i] = T[i] + 32
		}
	}

}

func cetakArray(T tabChar, n int) {
	/*I.S. Terdefinisi sekumpulan n karakter pada array T
	  F.S. Menampilkan seluruh elemen array T pada layar
	*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%c", T[i])
	}
}

func main() {
	var tab tabChar
	var n int
	masukanArray(&tab, &n)
	lowerCase(&tab, n)
	cetakArray(tab, n)
}
